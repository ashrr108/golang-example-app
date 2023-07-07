package main

import (
	"github.com/aristat/golang-example-app/app/entrypoint"
	"github.com/aristat/golang-example-app/app/http"
	"github.com/aristat/golang-example-app/app/logger"
	"github.com/spf13/cobra"
)

func main() {
	// Initialize the application
	app := entrypoint.New()

	// Add services to the application
	app.AddService(http.New())
	app.AddService(logger.New())

	// Add commands to the application
	app.AddCommand(cobra.Command{
		Use:   "run",
		Short: "Run the application",
		Run: func(cmd *cobra.Command, args []string) {
			app.Run()
		},
	})

	// Execute the application
	app.Execute()
}
