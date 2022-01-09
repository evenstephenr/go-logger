# go-logger

A sample go package that I stood up quickly to learn how remote modules work in go.

## Installing this package

From inside a go project...

1. run `go get github.com/evenstephenr/go-logger@core`
1. inside a `.go` file, just import the logger and then call it

```go
// $USERPROFILE/go/package
package main

import "github.com/evenstephenr/go-logger"

func main() {
  logger.Info("inside main package", map[string]string{ "ts": time.Now().String() })
}

```