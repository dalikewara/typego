package typego_test

import (
	"fmt"
	"github.com/dalikewara/typego"
	"log"
	"testing"
)

var defaultErrorLogHandler = func(err typego.Error) {
	log.Println(fmt.Sprintf("%+v", err))
}

var defaultInfoLogHandler = func(info typego.Info) {
	log.Println(fmt.Sprintf("%+v", info))
}

func TestSetCustomErrorLog(t *testing.T) {
	errGeneral := typego.NewError("01", "general error")

	_ = errGeneral.Log()

	typego.SetCustomErrorLog(func(err typego.Error) {
		fmt.Println(fmt.Sprintf("hello i am a custom log! -> %+v", err))
	})

	_ = errGeneral.Log()

	typego.SetCustomErrorLog(defaultErrorLogHandler)

	_ = errGeneral.Log()
}

func TestSetCustomInfoLog(t *testing.T) {
	infoGeneral := typego.NewInfo()

	_ = infoGeneral.Log()

	typego.SetCustomInfoLog(func(info typego.Info) {
		fmt.Println(fmt.Sprintf("hello i am a custom log! -> %+v", info))
	})

	_ = infoGeneral.Log()

	typego.SetCustomInfoLog(defaultInfoLogHandler)

	_ = infoGeneral.Log()
}
