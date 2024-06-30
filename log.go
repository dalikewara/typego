package typego

import (
	"fmt"
	"log"
)

var errorLogHandler = func(err Error) {
	log.Println(fmt.Sprintf("%+v", err))
}

type ErrorLogHandler func(err Error)

// SetCustomErrorLog sets custom error log handler
func SetCustomErrorLog(handler ErrorLogHandler) {
	errorLogHandler = handler
}
