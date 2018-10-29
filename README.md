[![Build Status](https://travis-ci.org/avence12/licenses.png?branch=master)](https://travis-ci.org/avence12/licenses)

`licenses` uses `go list` tool over a Go workspace to collect the dependencies
of a package or command, detect their license if any and match them against
well-known templates.

Inspired by [https://github.com/pmezard/licenses](https://github.com/pmezard/licenses).

## New Features

- Support Go Modules
- Update list of license
    - Add [Server Side Public License (SSPL)](https://www.mongodb.com/licensing/server-side-public-license)

## Quick Start

```
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

Unmatched license words can be displayed with:

```
$ licenses -w github.com/steveyen/gtreap
github.com/steveyen/gtreap  MIT License (98%)
                            -words: mit, license
```
