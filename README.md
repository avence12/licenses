[![Build Status](https://travis-ci.org/avence12/licenses.png?branch=master)](https://travis-ci.org/avence12/licenses)
[![Go Report Card](https://goreportcard.com/badge/github.com/avence12/licenses)](https://goreportcard.com/report/github.com/avence12/licenses)
[![GoDoc](https://godoc.org/github.com/avence12/licenses?status.svg)](https://godoc.org/github.com/avence12/licenses)
[![Release](https://img.shields.io/github/release/avence12/licenses.svg?style=flat-square)](https://github.com/avence12/licenses/releases)

`licenses` uses `go list` tool over a Go workspace to collect the dependencies
of a package or command, detect their license if any and match them against
well-known templates. Require **Golang 1.11** or above to build this project.

Inspired by [https://github.com/pmezard/licenses](https://github.com/pmezard/licenses).

## Installation

```sh
$ go get -u github.com/avence12/licenses
```

## Quick Start

- Show all licenses of `github.com/drone/drone-cli/drone` and its dependencies

```sh
# get source code
$ go get -u -d github.com/drone/drone-cli/drone


# check licenses
$ licenses github.com/drone/drone-cli/drone
github.com/drone/drone-cli                                                  Apache License 2.0
github.com/drone/drone-cli/vendor/github.com/Sirupsen/logrus                MIT License
github.com/drone/drone-cli/vendor/github.com/docker/distribution/reference  Apache License 2.0
github.com/drone/drone-cli/vendor/github.com/docker/docker                  Apache License 2.0 (96%)
github.com/drone/drone-cli/vendor/github.com/docker/go-connections          Apache License 2.0 (96%)
github.com/drone/drone-cli/vendor/github.com/docker/go-units                Apache License 2.0 (96%)
github.com/drone/drone-cli/vendor/github.com/docker/libcompose/yaml         Apache License 2.0 (96%)
github.com/drone/drone-cli/vendor/github.com/drone/drone-go/drone           Apache License 2.0
github.com/drone/drone-cli/vendor/github.com/drone/envsubst/parse           MIT License
github.com/drone/drone-cli/vendor/github.com/flynn/go-shlex                 Apache License 2.0
github.com/drone/drone-cli/vendor/github.com/ghodss/yaml                    ? (BSD 3-clause "New" or "Revised" License, 83%)
github.com/drone/drone-cli/vendor/github.com/jackspirou/syscerts            Apache License 2.0
github.com/drone/drone-cli/vendor/github.com/joho/godotenv/autoload         MIT License
github.com/drone/drone-cli/vendor/github.com/pkg/errors                     BSD 2-clause "Simplified" License
github.com/drone/drone-cli/vendor/github.com/urfave/cli                     MIT License
github.com/drone/drone-cli/vendor/golang.org/x/net                          BSD 3-clause "New" or "Revised" License (96%)
github.com/drone/drone-cli/vendor/golang.org/x/oauth2/internal              BSD 3-clause "New" or "Revised" License (96%)
github.com/drone/drone-cli/vendor/golang.org/x/sync/errgroup                BSD 3-clause "New" or "Revised" License (96%)
github.com/drone/drone-cli/vendor/gopkg.in/yaml.v2                          ? (The Unlicense, 35%)
```

- Unmatched license words can be displayed with:

```sh
$ licenses -w github.com/docker/go-units
github.com/docker/go-units  Apache License 2.0 (96%)
                            +words: https
                            -words: http, how, apply, attach, boilerplate, fields
```
