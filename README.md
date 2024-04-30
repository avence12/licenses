[![Build Status](https://travis-ci.org/avence12/licenses.png?branch=master)](https://travis-ci.org/avence12/licenses)
[![Go Report Card](https://goreportcard.com/badge/github.com/avence12/licenses)](https://goreportcard.com/report/github.com/avence12/licenses)
[![GoDoc](https://godoc.org/github.com/avence12/licenses?status.svg)](https://godoc.org/github.com/avence12/licenses)
[![Release](https://img.shields.io/github/release/avence12/licenses.svg?style=flat-square)](https://github.com/avence12/licenses/releases)

`licenses` uses `go list` tool over a Go workspace to collect the dependencies
of a package or command, detect their license if any and match them against
well-known templates. Require **Golang 1.16** or above to build this project such supports [Go Modules](https://go.dev/ref/mod).

Inspired by [https://github.com/pmezard/licenses](https://github.com/pmezard/licenses).

## Installation

```sh
$ go install github.com/avence12/licenses@latest
```

## Quick Start

- Show all licenses of `github.com/spf13/cobra` and its dependencies

```sh
# Download source code and get to the download folder
$ go get -u -d github.com/spf13/cobra
$ cd $GOPATH/pkg/mod/github.com/spf13/cobra*

# Check licenses
$ licenses github.com/spf13/cobra

github.com/inconshreveable/mousetrap  Apache License 2.0
github.com/spf13/cobra                Apache License 2.0 (95%)
github.com/spf13/pflag                BSD 3-clause "New" or "Revised" License (96%)

# Unmatched license words can be displayed with:
$ licenses -w github.com/spf13/pflag
github.com/spf13/pflag  BSD 3-clause "New" or "Revised" License (96%)
                        +words: google, inc, owner
                        -words: all, rights, reserved, project, holder
```
