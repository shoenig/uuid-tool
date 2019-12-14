uuid-tool
=========

Use `uuid-tool` to generate v4 UUIDs on the command line.

[![Go Report Card](https://goreportcard.com/badge/gophers.dev/cmds/uuid-tool)](https://goreportcard.com/report/gophers.dev/cmds/uuid-tool)
[![Build Status](https://travis-ci.org/shoenig/uuid-tool.svg?branch=master)](https://travis-ci.org/shoenig/uuid-tool)
[![GoDoc](https://godoc.org/gophers.dev/cmds/uuid-tool?status.svg)](https://godoc.org/gophers.dev/cmds/uuid-tool)
![NetflixOSS Lifecycle](https://img.shields.io/osslifecycle/shoenig/uuid-tool.svg)
![GitHub](https://img.shields.io/github/license/shoenig/uuid-tool.svg)

# Project Overview

Module `gophers.dev/cmds/uuid-tool` provides a command-line utility for generating
v4 UUIDs.

Under the hood, `uuid-tool` makes use of HashiCorp's [go-uuid](https://github.com/hashicorp/go-uuid)
library for v4 UUID generation.

# Getting Started

The `uuid-tool` command can be installed by running
```bash
$ go get gophers.dev/cmds/uuid-tool
```

# Example Usages

#### generate single v4 UUID with trailing newline (default)
```bash
$ uuid-tool
fafae80e-a308-120d-99fa-07ecd010e2b0
```

#### generate multiple v4 UUIDs (-c [count])
```bash
$ uuid-tool -c 3
b723aaef-0811-bbbc-7d6b-69555a683af8
11fa89a0-d0a4-e8d3-c89f-4d1907b7f385
b27ff498-83ae-eb29-263b-560e118942d7
```

#### omit trailing newline (-n)
```bash
# useful for pasting directly to clipboard
$ uuid-tool -n | pcbopy
```

# Contributing
The `gophers.dev/cmds/uuid-tool` module is always improving with new features
and error corrections. For contributing bug fixes and new features please file
an issue.

# License
The `gophers.dev/cmds/uuid-tool` module is open source under the [MIT](LICENSE) license.
