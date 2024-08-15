package main

import (
	"fmt"
	"github.com/r4ndsen/eancheckr"
	"os"
)

const (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"
)

func main() {
	if err := run(); err != nil {
		warn(err.Error())
		os.Exit(1)
	}

	info("valid")
}

func info(message string) {
	fmt.Fprintf(os.Stdout, "%s%s%s\n", Green, message, Reset)
}

func warn(message string) {
	fmt.Fprintf(os.Stderr, "%s%s%s\n", Red, message, Reset)
}

func run() error {
	if len(os.Args) != 2 {
		return fmt.Errorf("usage: %s %q", os.Args[0], "<ean>")
	}

	_, err := eancheckr.VerifyEan(os.Args[1])

	return err
}
