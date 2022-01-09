package logger

import "fmt"

func Error(message string, context interface{}) {
	fmt.Println("ERROR: ", message, context)
}

func Info(message string, context interface{}) {
	fmt.Println("INFO: ", message, context)
}