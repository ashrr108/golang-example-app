package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	packages := []string{
		"github.com/golang-jwt/jwt",
		"github.com/spf13/cast",
		"github.com/spf13/cobra",
	}

	for _, pkg := range packages {
		err := exec.Command("go", "get", pkg).Run()
		if err != nil {
			fmt.Printf("Failed to get package %s: %s\n", pkg, err)
			os.Exit(1)
		}
	}

	fmt.Println("Successfully fetched all packages")
}
