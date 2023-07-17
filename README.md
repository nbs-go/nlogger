# nlogger

[![Go Report Card](https://goreportcard.com/badge/github.com/nbs-go/nlogger)](https://goreportcard.com/report/github.com/nbs-go/nlogger)
[![GitHub license](https://img.shields.io/github/license/nbs-go/nlogger)](https://github.com/nbs-go/nlogger/blob/master/LICENSE)
[![codecov](https://codecov.io/gh/nbs-go/nlogger/branch/master/graph/badge.svg?token=ZPPD4U6JCE)](https://codecov.io/gh/nbs-go/nlogger)

A Logger interface for Golang. Inspired by `database/sql` package. Batteries included.

## Install

```shell
go get github.com/nbs-go/nlogger
```

## Usage

```
package main

import (
  // Register a logger implementation, just do it once in the main package
  // Make sure to place this import at the first line, in the first files of main package (ordered by alphabet)
  _ "github.com/nbs-go/nlogger-json"

  "github.com/nbs-go/nlogger"
)

func main() {
  // If no logger registered, it will use nlogger.StdLogger which are
  // based on built-in go "log" package
  log := nlogger.Get()

  log.Info("Hello World")
}
```

## TODO

- [ ] Documentation

## Contributors

- Saggaf Arsyad <saggaf.arsyad@gmail.com>
