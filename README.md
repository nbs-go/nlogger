# nlogger

[![Go Report Card](https://goreportcard.com/badge/github.com/nbs-go/nlogger)](https://goreportcard.com/report/github.com/nbs-go/nlogger)
[![GitHub license](https://img.shields.io/github/license/nbs-go/nlogger)](https://github.com/nbs-go/nlogger/blob/master/LICENSE)
[![codecov](https://codecov.io/gh/nbs-go/nlogger/branch/master/graph/badge.svg?token=ZPPD4U6JCE)](https://codecov.io/gh/nbs-go/nlogger)

A Logger interface for Golang. Inspired by `database/sql` package. Batteries included.

## Installing

```shell
go get github.com/nbs-go/nlogger
```

## Example

```
package main

import (
  "github.com/nbs-go/nlogger"
  
  // Register a logger implementation, just do it once in the main package
  // _ "github.com/nbs-go/nlogrus"
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
