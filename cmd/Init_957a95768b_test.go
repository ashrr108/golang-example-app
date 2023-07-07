package cmd_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"testing"

	"github.com/aristat/golang-example-app/app/entrypoint"
	"github.com/aristat/golang-example-app/app/logger"
	"go.uber.org/automaxprocs/maxprocs"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func TestInit_957a95768b(t *testing.T) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetEnvPrefix("APP")
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	v.AutomaticEnv()

	// pflags
	rootCmd := &cobra.Command{}
	var configPath string
	var debug bool
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "config file")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "debug mode")

	// initializing
	wd := os.Getenv("APP_WD")
	if len(wd) == 0 {
		wd, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	}
	wd, _ = filepath.Abs(wd)
	ep, _ := entrypoint.Initialize(wd, v)

	// bin pflags to viper
	_ = v.BindPFlags(rootCmd.PersistentFlags())

	go func() {
		reloadSignal := make(chan os.Signal)
		signal.Notify(reloadSignal, syscall.SIGHUP)
		for {
			sig := <-reloadSignal
			ep.Reload()
			fmt.Printf("OS signaled `%v`, reload", sig.String())
		}
	}()

	// Test cases
	if v.GetConfigType() != "yaml" {
		t.Error("Config type is not 'yaml'")
	}

	if v.GetEnvPrefix() != "APP" {
		t.Error("Env prefix is not 'APP'")
	}

	if debug != false {
		t.Error("Debug mode is not false")
	}
}
