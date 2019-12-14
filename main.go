package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(name(os.Args[0]), os.Args[1:], os.Stdout); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
