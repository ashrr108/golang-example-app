package command

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/aristat/golang-example-app/command/daemon"
	"github.com/aristat/golang-example-app/command/health_check_service"
	"github.com/aristat/golang-example-app/command/jwt"
	"github.com/aristat/golang-example-app/command/migrate"
	"github.com/aristat/golang-example-app/command/product_service"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "Root command",
}

func Execute() error {
	rootCmd.AddCommand(daemon.Cmd, product_service.Cmd, health_check_service.Cmd, migrate.Cmd, jwt.Cmd)
	return rootCmd.Execute()
}

func TestExecute_0263ba38e4(t *testing.T) {
	err := Execute()
	if err != nil {
		t.Error("Expected successful execution, got error", err)
	}

	oldStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("Execution error")
	}

	err = Execute()

	w.Close()
	os.Stderr = oldStderr

	var buf bytes.Buffer
	buf.ReadFrom(r)
	errMsg := buf.String()

	if err == nil || errMsg != "Execution error\n" {
		t.Error("Expected 'Execution error', got", errMsg)
	}
}
