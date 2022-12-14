package main

import (
	"context"
	"fmt"
	"github.com/btcsuite/btcwallet/walletdb"
	"github.com/ellemouton/sqlite/common_sql"
	"log"
	_ "modernc.org/sqlite"
	"net/url"
	"sync"
	"time"
)

const (
	// sqliteOptionPrefix is the string prefix sqlite uses to set various
	// options. This is used in the following format:
	//   * sqliteOptionPrefix || option_name = option_value.
	sqliteOptionPrefix = "_pragma"
)

// postgresReplacements define a set of postgres keywords that should be swapped
// out with certain other sqlite keywords in any queries.
var postgresReplacements = common_sql.PostgresCmdReplacements{
	"BYTEA":                 "BLOB",
	"BIGSERIAL PRIMARY KEY": "INTEGER PRIMARY KEY AUTOINCREMENT",
}

func main() {
	if err := doThings(); err != nil {
		log.Fatalln(err)
	}
}

func doThings() error {
	// The set of pragma options are accepted using query options. For now
	// we only want to ensure that foreign key constraints are properly
	// enforced.
	pragmaOptions := []struct {
		name  string
		value string
	}{
		{
			name:  "busy_timeout",
			value: "500",
		},
		{
			name:  "foreign_keys",
			value: "on",
		},
		{
			name:  "journal_mode",
			value: "WAL",
		},
	}
	sqliteOptions := make(url.Values)
	for _, option := range pragmaOptions {
		sqliteOptions.Add(
			sqliteOptionPrefix,
			fmt.Sprintf("%v=%v", option.name, option.value),
		)
	}

	// Construct the DSN which is just the database file name, appended
	// with the series of pragma options as a query URL string. For more
	// details on the formatting here, see the modernc.org/sqlite docs:
	// https://pkg.go.dev/modernc.org/sqlite#Driver.Open.
	dsn := fmt.Sprintf(
		"%v?%v", "sqlite-test", sqliteOptions.Encode(),
	)

	fmt.Println(dsn)

	common_sql.Init(0)
	db, err := common_sql.NewSqlBackend(context.Background(), &common_sql.Config{
		DriverName:              "sqlite",
		Dsn:                     dsn,
		TableNamePrefix:         "testtable",
		PostgresCmdReplacements: postgresReplacements,
	})
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		err = db.Update(func(tx walletdb.ReadWriteTx) error {
			fmt.Println("WE ARE HERE")
			_, err := tx.CreateTopLevelBucket([]byte{1, 1, 1, 1})

			time.Sleep(time.Second * 10)

			_, err = tx.CreateTopLevelBucket([]byte{2, 2, 2, 2})
			return err
		}, func() {})
		if err != nil {
			log.Fatalln(err)
		}
	}()

	//time.Sleep(time.Second)
	//fmt.Println("trying to get in...")
	//err = db.View(func(tx walletdb.ReadTx) error {
	//	fmt.Println("WE ARE IN!")
	//	tx.ReadBucket([]byte{1, 1, 1, 1})
	//	return nil
	//
	//}, func() {})
	//if err != nil {
	//	log.Fatalln(err)
	//}

	time.Sleep(time.Second)
	fmt.Println("trying to get in...")
	err = db.Update(func(tx walletdb.ReadWriteTx) error {
		fmt.Println("WE ARE IN!")
		_, err := tx.CreateTopLevelBucket([]byte{1, 1, 1, 1})
		return err
	}, func() {})
	if err != nil {
		log.Fatalln(err)
	}

	wg.Wait()

	return err
}
