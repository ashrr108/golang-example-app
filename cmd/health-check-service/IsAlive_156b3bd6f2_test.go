// First, install all the required packages
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	packages := []string{
		"github.com/aristat/golang-example-app/app/common",
		"github.com/aristat/golang-example-app/app/config",
		"github.com/aristat/golang-example-app/app/logger",
		"github.com/aristat/golang-example-app/generated/resources/proto/health_checks",
		"github.com/spf13/cobra",
		"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc",
		"go.opentelemetry.io/otel",
		"go.opentelemetry.io/otel/propagation",
		"google.golang.org/grpc",
		"google.golang.org/protobuf/types/known/emptypb",
	}

	for _, pkg := range packages {
		cmd := exec.Command("go", "get", pkg)
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Failed to get package %s: %v\n", pkg, err)
		}
	}
}
