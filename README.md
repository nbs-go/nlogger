# nlogger

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
