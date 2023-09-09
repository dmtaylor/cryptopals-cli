//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/target"
)

// Build builds application binary
func Build() error {
	toUpdate, err := target.Dir("cryptopals-cli", "main.go", "cmd", "util")
	if err != nil {
		return fmt.Errorf("failed to get update deps: %w", err)
	}
	if toUpdate {
		fmt.Printf("go build .\n")
		cmd := exec.Command("go", "build", ".")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}
	return nil
}

// Clean removes the application binary
func Clean() error {
	cmd := exec.Command("go", "clean", "-x")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
	return nil
}

// Rebuild rebuilds app from scratch
func Rebuild() error {
	mg.SerialDeps(Clean, Build)
	return nil
}

// Tests run all tests
func Tests() error {
	mg.Deps(Build)
	fmt.Println("running tests...")
	cmd := exec.Command("go", "test", "-v", "-bench=.", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
