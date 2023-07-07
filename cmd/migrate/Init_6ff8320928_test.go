package migrate

import (
	"testing"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/cobra"
)

var TestCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Manage database migrations",
}

var (
	testArgTable string
	testArgLimit int
	testArgDsn   string
)

func testInit() {
	TestCmd.PersistentFlags().StringVar(&testArgTable, "table", "migrations", "Name of the table in the database")
	TestCmd.PersistentFlags().IntVar(&testArgLimit, "limit", 1, "Max number of migrations to apply")
	TestCmd.PersistentFlags().StringVar(&testArgDsn, "dsn", "postgres://localhost:5432/golang_example_development?sslmode=disable", "Database connection string")
}

func TestInit_6ff8320928(t *testing.T) {
	pflag.VisitAll(func(f *pflag.Flag) {
		f.Value.Set(f.DefValue)
	})

	t.Run("default values", func(t *testing.T) {
		testInit()

		if testArgTable != "migrations" {
			t.Errorf("Expected default table to be 'migrations', got %s", testArgTable)
		}

		if testArgLimit != 1 {
			t.Errorf("Expected default limit to be 1, got %d", testArgLimit)
		}

		if testArgDsn != "postgres://localhost:5432/golang_example_development?sslmode=disable" {
			t.Errorf("Expected default dsn to be 'postgres://localhost:5432/golang_example_development?sslmode=disable', got %s", testArgDsn)
		}
	})

	t.Run("custom values", func(t *testing.T) {
		os.Args = []string{"cmd", "--table=custom_table", "--limit=5", "--dsn=custom_dsn"}
		testInit()

		if testArgTable != "custom_table" {
			t.Errorf("Expected table to be 'custom_table', got %s", testArgTable)
		}

		if testArgLimit != 5 {
			t.Errorf("Expected limit to be 5, got %d", testArgLimit)
		}

		if testArgDsn != "custom_dsn" {
			t.Errorf("Expected dsn to be 'custom_dsn', got %s", testArgDsn)
		}
	})
}
