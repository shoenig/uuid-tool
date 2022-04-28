package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/shoenig/test"
)

func Test_name(t *testing.T) {
	t.Parallel()

	try := func(t *testing.T, input, exp string) {
		result := name(input)
		test.EqCmp(t, exp, result)
	}

	t.Run("relative path", func(t *testing.T) {
		try(t, "./foo", "foo")
	})

	t.Run("full path", func(t *testing.T) {
		try(t, "/foo/bar/baz", "baz")
	})
}

func Test_generateOne(t *testing.T) {
	t.Parallel()

	try := func(t *testing.T, newline bool) {
		var buf bytes.Buffer
		err := generateOne(&buf, newline)
		test.NoError(t, err)
		result := buf.String()
		test.EqCmp(t, 37, len(result))
		endsNewline := strings.HasSuffix(result, "\n")
		test.EqCmp(t, newline, endsNewline)
	}

	t.Run("with newline", func(t *testing.T) {
		try(t, true)
	})

	t.Run("without newline", func(t *testing.T) {
		try(t, false)
	})
}

func Test_generate(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	err := generate(&buf, 10, true)
	test.NoError(t, err)
	result := buf.String()
	test.EqCmp(t, 37*10, len(result))
}
