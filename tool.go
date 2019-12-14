package main

import (
	"flag"
	"fmt"
	"io"
	"path/filepath"

	uuid "github.com/hashicorp/go-uuid"
	"github.com/pkg/errors"
)

func name(file string) string {
	return filepath.Base(file)
}

func run(cmd string, args []string, w io.Writer) error {
	fs := flag.NewFlagSet(cmd, flag.ExitOnError)
	nl := fs.Bool("n", false, "do not print trailing newline")
	c := fs.Int("c", 1, "the number of UUIDs to generate")
	if err := fs.Parse(args); err != nil {
		return err
	}
	return generate(w, *c, !*nl)
}

func generate(w io.Writer, c int, newline bool) error {
	for i := 0; i < c; i++ {
		if err := generateOne(w, newline); err != nil {
			return err
		}
	}
	return nil
}

func generateOne(w io.Writer, newline bool) error {
	id, err := uuid.GenerateUUID()
	if err != nil {
		return errors.Wrap(err, "unable to generate UUID")
	}

	lineFmt := "%s "
	if newline {
		lineFmt = "%s\n"
	}

	_, err = fmt.Fprintf(w, lineFmt, id)
	return err
}
