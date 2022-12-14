package common_sql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/btcsuite/btcwallet/walletdb"
)

const (
	// kvTableName is the name of the table that will contain all the kv
	// pairs.
	kvTableName = "kv"
)

// Config holds a set of configuration options of a sql database connection.
type Config struct {
	// DriverName is the string that defines the registered sql driver that
	// is to be used.
	DriverName string

	// Dsn is the database connection string that will be used to connect
	// to the db.
	Dsn string

	// Timeout is the time after which a query to the db will be canceled if
	// it has not yet completed.
	Timeout time.Duration

	// Schema is the name of the schema under which the sql tables should be
	// created. It should be left empty for backends like sqlite that do not
	// support having more than one schema.
	Schema string

	// TableNamePrefix is the name that should be used as a table name
	// prefix when constructing the KV style table.
	TableNamePrefix string

	// PostgresCmdReplacements define a one-to-one string mapping of
	// postgres keywords to the strings that should replace those postgres
	// keywords in any commands. If left empty, then postgres syntax will be
	// used in any sql queries.
	PostgresCmdReplacements PostgresCmdReplacements
}

// db holds a reference to the sql db connection.
type db struct {
	// cfg is the sql db connection config.
	cfg *Config

	// prefix is the table name prefix that is used to simulate namespaces.
	// We don't use schemas because at least sqlite does not support that.
	prefix string

	// ctx is the overall context for the database driver.
	//
	// TODO: This is an anti-pattern that is in place until the kvdb
	// interface supports a context.
	ctx context.Context

	// db is the underlying database connection instance.
	db *sql.DB

	// lock is the global write lock that ensures single writer.
	lock sync.RWMutex

	// table is the name of the table that contains the data for all
	// top-level buckets that have keys that cannot be mapped to a distinct
	// sql table.
	table string
}

// Enforce db implements the walletdb.DB interface.
var _ walletdb.DB = (*db)(nil)

// Global set of database connections.
var dbConns *dbConnSet

// Init initializes the global set of database connections.
func Init(maxConnections int) {
	dbConns = newDbConnSet(maxConnections)
}

// NewSqlBackend returns a db object initialized with the passed backend
// config. If database connection cannot be established, then returns error.
func NewSqlBackend(ctx context.Context, cfg *Config) (*db, error) {
	if dbConns == nil {
		return nil, errors.New("db connection set not initialized")
	}

	if cfg.TableNamePrefix == "" {
		return nil, errors.New("empty table name prefix")
	}

	table := fmt.Sprintf("%s_%s", cfg.TableNamePrefix, kvTableName)

	query := newKVSchemaCreationCmd(
		table, cfg.Schema, cfg.PostgresCmdReplacements,
	)

	dbConn, err := dbConns.Open(cfg.DriverName, cfg.Dsn)
	if err != nil {
		return nil, err
	}

	_, err = dbConn.ExecContext(ctx, query)
	if err != nil {
		_ = dbConn.Close()

		return nil, err
	}

	return &db{
		cfg:    cfg,
		ctx:    ctx,
		db:     dbConn,
		table:  table,
		prefix: cfg.TableNamePrefix,
	}, nil
}

// getTimeoutCtx gets a timeout context for database requests.
func (db *db) getTimeoutCtx() (context.Context, func()) {
	if db.cfg.Timeout == time.Duration(0) {
		return db.ctx, func() {}
	}

	return context.WithTimeout(db.ctx, db.cfg.Timeout)
}

// getPrefixedTableName returns a table name for this prefix (namespace).
func (db *db) getPrefixedTableName(table string) string {
	return fmt.Sprintf("%s_%s", db.prefix, table)
}

// catchPanic executes the specified function. If a panic occurs, it is returned
// as an error value.
func catchPanic(f func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Criticalf("Caught unhandled error: %v", r)

			switch data := r.(type) {
			case error:
				err = data

			default:
				err = errors.New(fmt.Sprintf("%v", data))
			}
		}
	}()

	err = f()

	return
}

// View opens a database read transaction and executes the function f with the
// transaction passed as a parameter. After f exits, the transaction is rolled
// back. If f errors, its error is returned, not a rollback error (if any
// occur). The passed reset function is called before the start of the
// transaction and can be used to reset intermediate state. As callers may
// expect retries of the f closure (depending on the database backend used), the
// reset function will be called before each retry respectively.
func (db *db) View(f func(tx walletdb.ReadTx) error, reset func()) error {
	return db.executeTransaction(
		func(tx walletdb.ReadWriteTx) error {
			return f(tx.(walletdb.ReadTx))
		},
		reset, true,
	)
}

// Update opens a database read/write transaction and executes the function f
// with the transaction passed as a parameter. After f exits, if f did not
// error, the transaction is committed. Otherwise, if f did error, the
// transaction is rolled back. If the rollback fails, the original error
// returned by f is still returned. If the commit fails, the commit error is
// returned. As callers may expect retries of the f closure, the reset function
// will be called before each retry respectively.
func (db *db) Update(f func(tx walletdb.ReadWriteTx) error, reset func()) (err error) {
	return db.executeTransaction(f, reset, false)
}

// executeTransaction creates a new read-only or read-write transaction and
// executes the given function within it.
func (db *db) executeTransaction(f func(tx walletdb.ReadWriteTx) error,
	reset func(), readOnly bool) error {

	reset()

	tx, err := newReadWriteTx(db, readOnly)
	if err != nil {
		return err
	}

	err = catchPanic(func() error { return f(tx) })
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Errorf("Error rolling back tx: %v", rollbackErr)
		}

		return err
	}

	return tx.Commit()
}

// PrintStats returns all collected stats pretty printed into a string.
func (db *db) PrintStats() string {
	return "stats not supported by Postgres driver"
}

// BeginReadWriteTx opens a database read+write transaction.
func (db *db) BeginReadWriteTx() (walletdb.ReadWriteTx, error) {
	return newReadWriteTx(db, false)
}

// BeginReadTx opens a database read transaction.
func (db *db) BeginReadTx() (walletdb.ReadTx, error) {
	return newReadWriteTx(db, true)
}

// Copy writes a copy of the database to the provided writer. This call will
// start a read-only transaction to perform all operations.
// This function is part of the walletdb.Db interface implementation.
func (db *db) Copy(w io.Writer) error {
	return errors.New("not implemented")
}

// Close cleanly shuts down the database and syncs all data.
// This function is part of the walletdb.Db interface implementation.
func (db *db) Close() error {
	log.Infof("Closing database %v", db.prefix)

	return dbConns.Close(db.cfg.Dsn)
}
