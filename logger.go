package logger

import (
	"flag"
	"fmt"
)

var debug = flag.Bool("debug", false, "enable debug logs")
var debugShort = flag.Bool("d", false, "enable debug logs")

func Error(message string) string {
	return ("ERROR: " + message)
}

func Info(message string) string {
	return ("INFO: " + message)
}

func Warn(message string) string {
	return ("WARN: " + message)
}

func Debug(message string) string {
	if (*debug) || (*debugShort) {
		return ("DEBUG: " + message)
	}

	return ""
}

func ReportContext(s string, context interface{}) {
	if (len(s) > 0) {
		fmt.Println(s, context)
	}
}

func Report(s string) {
	if (len(s) > 0) {
		fmt.Println(s)
	}
}