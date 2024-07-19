package typego

import (
	"fmt"
	"log"
)

var errorLogHandler = func(err Error) {
	log.Println(fmt.Sprintf("%+v", err))
}

var infoLogHandler = func(info Info) {
	log.Println(fmt.Sprintf("info: %+v", info))
}

type ErrorLogHandler func(err Error)

type InfoLogHandler func(info Info)

// SetCustomErrorLog sets custom error log handler
func SetCustomErrorLog(handler ErrorLogHandler) {
	errorLogHandler = handler
}

// SetCustomInfoLog sets custom info log handler
func SetCustomInfoLog(handler InfoLogHandler) {
	infoLogHandler = handler
}
