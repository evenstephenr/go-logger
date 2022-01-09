package logger

import (
	"flag"
	"fmt"
)

var debug = flag.Bool("debug", false, "enable debug logs")
var debugShort = flag.Bool("d", false, "enable debug logs")

func Error(message string, context interface{}) {
	fmt.Println("ERROR:", message, context)
}

func Info(message string, context interface{}) {
	fmt.Println("INFO:", message, context)
}

func Warn(message string, context interface{}) {
	fmt.Println("WARN:", message, context)
}

func Debug(message string, context interface{}) {
	if (*debug) || (*debugShort) {
		fmt.Println("DEBUG:", message, context)	
	}
}