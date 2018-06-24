package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func run(t *testing.T, args ...string) error {
	fmt.Printf("$ %s\n", strings.Join(args, " "))
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func TestCmdLine(t *testing.T) {
	if err := run(t, "go", "build"); err != nil {
		t.Fatalf("Can't build - %s", err)
	}

	if err := run(t, "docker", "build", "-f", "Dockerfile.test", "."); err != nil {
		t.Fatal()
	}
}
