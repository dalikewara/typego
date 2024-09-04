package typego

import (
	"encoding/json"
	"fmt"
)

var errorLogHandler = func(err Error) {
	b, e := json.Marshal(err)
	if e != nil {
		fmt.Println(e)
	}

	fmt.Println(string(b))
}

var infoLogHandler = func(info Info) {
	b, e := json.Marshal(info)
	if e != nil {
		fmt.Println(e)
	}

	fmt.Println(string(b))
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
