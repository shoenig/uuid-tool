uuid-tool
=========

Use `uuid-tool` to generate v4 UUIDs on the command line.

![GitHub](https://img.shields.io/github/license/shoenig/uuid-tool.svg)


[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-white.svg)](https://snapcraft.io/uuid-tool)

# Project Overview

Module `github.com/shoenig/uuid-tool` provides a command-line utility for generating
v4 UUIDs.

Under the hood, `uuid-tool` makes use of HashiCorp's [go-uuid](https://github.com/hashicorp/go-uuid)
library for v4 UUID generation.

# Getting Started

#### Install from Releases

The `uuid-tool` command can be downloaded from the [Releases](https://github.com/shoenig/uuid-tool/releases) page.

#### Install from SnapCraft
The `uuid-tool` command can be installed as a snap
```bash
$ sudo snap install uuid-tool
```

#### Build from source
The `uuid-tool` command can be compiled by running
```bash
$ go install github.com/shoenig/uuid-tool@latest
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
The `github.com/shoenig/uuid-tool` module is always improving with new features
and error corrections. For contributing bug fixes and new features please file
an issue.

# License
The `github.com/shoenig/uuid-tool` module is open source under the [MIT](LICENSE) license.
