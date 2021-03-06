// +build cockroachdb

package testfixtures

import (
	"os"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

func TestCockroachDB(t *testing.T) {
	for _, driver := range []string{"postgres", "pgx"} {
		testLoader(
			t,
			"cockroachdb",
			driver,
			os.Getenv("CRDB_CONN_STRING"),
			"testdata/schema/postgresql.sql",
			DangerousSkipTestDatabaseCheck(),
		)
	}
}
