package main

import (
	"fmt"
	"os"

	"github.com/tpology/auth"
)

func main() {
	if err := auth.New().Command().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
