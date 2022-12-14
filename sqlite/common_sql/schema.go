package common_sql

import (
	"fmt"
	"strings"
)

type PostgresCmdReplacements map[string]string

func newKVSchemaCreationCmd(table, schema string,
	replacements PostgresCmdReplacements) string {

	var (
		tableInSchema = table
		finalCmd      string
	)
	if schema != "" {
		finalCmd = fmt.Sprintf(
			`CREATE SCHEMA IF NOT EXISTS ` + schema + `;`,
		)

		tableInSchema = fmt.Sprintf("%s.%s", schema, table)
	}

	// Construct the sql statements to set up a kv table in postgres. Every
	// row points to the bucket that it is one via its parent_id field. A
	// NULL parent_id means that the key belongs to the upper-most bucket in
	// this table. A constraint on parent_id is enforcing referential
	// integrity.
	//
	// Furthermore there is a <table>_p index on parent_id that is required
	// for the foreign key constraint.
	//
	// Finally there are unique indices on (parent_id, key) to prevent the
	// same key being present in a bucket more than once (<table>_up and
	// <table>_unp). In postgres, a single index wouldn't enforce the unique
	// constraint on rows with a NULL parent_id. Therefore two indices are
	// defined.
	finalCmd += fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS ` + tableInSchema + `
(
    key BYTEA NOT NULL,
    value BYTEA,
    parent_id BIGINT,
    id BIGSERIAL PRIMARY KEY,
    sequence BIGINT,
    CONSTRAINT ` + table + `_parent FOREIGN KEY (parent_id)
        REFERENCES ` + tableInSchema + ` (id)
        ON UPDATE NO ACTION
        ON DELETE CASCADE
);
CREATE INDEX IF NOT EXISTS ` + table + `_p
    ON ` + tableInSchema + ` (parent_id);
CREATE UNIQUE INDEX IF NOT EXISTS ` + table + `_up
    ON ` + tableInSchema + `
    (parent_id, key) WHERE parent_id IS NOT NULL;
CREATE UNIQUE INDEX IF NOT EXISTS ` + table + `_unp 
    ON ` + tableInSchema + ` (key) WHERE parent_id IS NULL;
`)

	for from, to := range replacements {
		finalCmd = strings.Replace(finalCmd, from, to, -1)
	}

	return finalCmd
}
