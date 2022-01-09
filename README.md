# go-logger

A sample go module that I stood up quickly to learn how remote modules work in go.

## Installing this module

From inside a go project,

1. run `go get github.com/evenstephenr/go-logger@core`
1. inside a `.go` file, just import the logger and then call it

```go
// $USERPROFILE/go/hello
package main

import "github.com/evenstephenr/go-logger"

func main() {
  logger.Info("inside main package", map[string]string{ "ts": time.Now().String() })
}

```

## Using this module

There are generic error levels that this logger can report - `Error`, `Info`, `Warn`. 

All generic errors generate strings with the error level prefixed.

```go
func Error(message string) string {
	return ("ERROR: " + message)
}

func Info(message string) string {
	return ("INFO: " + message)
}

func Warn(message string) string {
	return ("WARN: " + message)
}
```

You can log these strings to the terminal yourself or use the built-in reporter.

```go
package main

import (
  "fmt"

  "github.com/evenstephenr/go-logger"
)

func main() {
  fmt.PrintLn(logger.Info("inside some function block..."))
  logger.Report(logger.Info("inside some function block..."))
}
```

There's a separate reporter that can log a map of string metadata inline.

```go
package main

import (
  "fmt"

  "github.com/evenstephenr/go-logger"
)

func main() {
	logger.ReportContext(logger.Debug("inside some function block"), map[string]string {
    "timestamp": time.Now().String()
	})
}
```

## Debug logs

The `Debug` error is a special error in this module. `Debug` errors will not be reported unless you add a `-d` or `-debug` flag to a go package invocation. If you are using your own reporter the `Debug` function returns an empty string when the package runtime does not include a debug flag.


```go
// $USERPROFILE/go/hello
package main

import (
  "fmt"
  
  "github.com/evenstephenr/go-logger"
)

func main() {
  fmt.PrintLn("Hello")
  logger.Report(logger.Debug("Hello Debug!")
  
}
```

```
$ go run .
Hello

$ go run . -d
Hello
Hello Debug!
```