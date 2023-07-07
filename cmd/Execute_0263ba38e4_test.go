package cmd

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
)

func TestExecute_0263ba38e4(t *testing.T) {
	rootCmd := &cobra.Command{
		Use:   "root",
		Short: "Root command",
		Long:  `This is the root command.`,
	}

	t.Run("success", func(t *testing.T) {
		testCmd := &cobra.Command{
			Use:   "test",
			Short: "Test command",
			Long:  `This is a test command.`,
		}
		rootCmd.AddCommand(testCmd)

		b := bytes.NewBufferString("")
		rootCmd.SetOut(b)
		rootCmd.SetArgs([]string{"test"})

		if err := rootCmd.Execute(); err != nil {
			t.Error(err)
		}

		out := b.String()
		if out != "Test command\n" {
			t.Errorf("Expected 'Test command', got '%s'", out)
		}
	})

	t.Run("failure", func(t *testing.T) {
		testCmd := &cobra.Command{
			Use:   "test",
			Short: "Test command",
			Long:  `This is a test command.`,
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Println("This is a failure test")
				os.Exit(1)
			},
		}
		rootCmd.AddCommand(testCmd)

		b := bytes.NewBufferString("")
		rootCmd.SetOut(b)
		rootCmd.SetArgs([]string{"test"})

		if err := rootCmd.Execute(); err == nil {
			t.Error("Expected error, got nil")
		}
	})
}
