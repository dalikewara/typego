package typego_test

import (
	"errors"
	"fmt"
	"github.com/dalikewara/typego"
	"log"
	"sync"
	"testing"
	"time"
)

func TestNewError(t *testing.T) {
	if err := typego.NewError("", ""); err == nil {
		log.Fatal("`err` must not nil")
	}

	if errGeneral := typego.NewError("01", "general error"); errGeneral == nil {
		log.Fatal("`errGeneral` must not nil")
	}
}

func TestErrorModel_ChangeCode(t *testing.T) {
	if err := typego.NewError("", "").ChangeCode("01"); err == nil {
		log.Fatal("`err` must not nil")
	}
}

func TestErrorModel_ChangeMessage(t *testing.T) {
	if err := typego.NewError("", "").ChangeMessage("general error"); err == nil {
		log.Fatal("`err` must not nil")
	}
}

func TestErrorModel_AddInfo(t *testing.T) {
	if err := typego.NewError("", "").AddInfo(errors.New("raw error")); err == nil {
		log.Fatal("`err` must not nil")
	}
}

func TestErrorModel_SetHttpStatus(t *testing.T) {
	if err := typego.NewError("", "").SetHttpStatus(500); err == nil {
		log.Fatal("`err` must not nil")
	}
}

func TestErrorModel_SetRPCStatus(t *testing.T) {
	if err := typego.NewError("", "").SetRPCStatus(13); err == nil {
		log.Fatal("`err` must not nil")
	}
}

func TestErrorModel_Log(t *testing.T) {
	_ = typego.NewError("01", "general error").Log()
}

func TestErrorModel_GetCode(t *testing.T) {
	if errCode := typego.NewError("01", "").GetCode(); errCode != "01" {
		log.Fatal("`errCode` must be `01`")
	}

	if errCode := typego.NewError("01", "").ChangeCode("02").GetCode(); errCode != "02" {
		log.Fatal("`errCode` must be `02`")
	}
}

func TestErrorModel_GetMessage(t *testing.T) {
	if errMessage := typego.NewError("", "general error").GetMessage(); errMessage != "general error" {
		log.Fatal("`errMessage` must be `general error`")
	}

	if errMessage := typego.NewError("", "general error").ChangeMessage("general error 2").GetMessage(); errMessage != "general error 2" {
		log.Fatal("`errMessage` must be `general error 2`")
	}
}

func TestErrorModel_GetInfo(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		if errInfos := typego.NewError("", "").AddInfo(errors.New("raw error"), errors.New("raw error 2")).AddInfo(errors.New("raw error 3")).GetInfo(); fmt.Sprintf("%v", errInfos) != fmt.Sprintf("%v", []string{"raw error", "raw error 2", "raw error 3"}) {
			log.Fatal("`errInfos` must be `[]string{\"raw error\", \"raw error 2\", \"raw error 3\"}`")
		}
	})

	t.Run("string", func(t *testing.T) {
		if errInfos := typego.NewError("", "").AddInfo("raw error", "raw error 2").AddInfo("raw error 3").GetInfo(); fmt.Sprintf("%v", errInfos) != fmt.Sprintf("%v", []string{"raw error", "raw error 2", "raw error 3"}) {
			log.Fatal("`errInfos` must be `[]string{\"raw error\", \"raw error 2\", \"raw error 3\"}`")
		}
	})

	t.Run("any", func(t *testing.T) {
		if errInfos := typego.NewError("", "").AddInfo(1).GetInfo(); fmt.Sprintf("%v", errInfos) != fmt.Sprintf("%v", []string{"1"}) {
			log.Fatal("`errInfos` must be `[]string{\"1\"}`")
		}
	})
}

func TestErrorModel_GetHttpStatus(t *testing.T) {
	if errHttpStatus := typego.NewError("", "").GetHttpStatus(); errHttpStatus != 0 {
		log.Fatal("`errHttpStatus` must be `0`")
	}

	if errHttpStatus := typego.NewError("", "").SetHttpStatus(404).GetHttpStatus(); errHttpStatus != 404 {
		log.Fatal("`errHttpStatus` must be `404`")
	}
}

func TestErrorModel_GetRPCStatus(t *testing.T) {
	if errRPCStatus := typego.NewError("", "").GetRPCStatus(); errRPCStatus != 0 {
		log.Fatal("`errRPCStatus` must be `0`")
	}

	if errRPCStatus := typego.NewError("", "").SetRPCStatus(10).GetRPCStatus(); errRPCStatus != 10 {
		log.Fatal("`errRPCStatus` must be `10`")
	}
}

func TestErrorModel_Error(t *testing.T) {
	if err := typego.NewError("01", "general error").SetHttpStatus(500).SetRPCStatus(13).AddInfo(errors.New("raw error").Error()).AddInfo("raw error 2").Error(); err != "error: {\"code\":\"01\",\"message\":\"general error\",\"info\":[\"raw error\",\"raw error 2\"],\"http_status\":500,\"rpc_status\":13}" {
		log.Fatal("`err` must be `error: {\"code\":\"01\",\"message\":\"general error\",\"info\":[\"raw error\",\"raw error 2\"],\"http_status\":500,\"rpc_status\":13}`")
	}
}

func TestNewErrorFromError(t *testing.T) {
	t.Run("valid_format", func(t *testing.T) {
		err := typego.NewErrorFromError(errors.New("error: {\"code\":\"01\",\"message\":\"general error\",\"http_status\":500,\"info\":[\"raw info\",\"raw info 2\"],\"rpc_status\":13}"))

		if errCode := err.GetCode(); errCode != "01" {
			log.Fatal("`errCode` must be `01`")
		}

		if errMessage := err.GetMessage(); errMessage != "general error" {
			log.Fatal("`errMessage` must be `general error`")
		}

		if errHttpStatus := err.GetHttpStatus(); errHttpStatus != 500 {
			log.Fatal("`errHttpStatus` must be `500`")
		}

		if errRPCStatus := err.GetRPCStatus(); errRPCStatus != 13 {
			log.Fatal("`errRPCStatus` must be `13`")
		}

		if errInfoLen := len(err.GetInfo()); errInfoLen != 2 {
			log.Fatal("`errInfoLen` must be `2`")
		}

		if errInfo1 := err.GetInfo()[0]; errInfo1 != "raw info" {
			log.Fatal("`errInfo1` must be `raw info`")
		}

		if errInfo2 := err.GetInfo()[1]; errInfo2 != "raw info 2" {
			log.Fatal("`errInfo2` must be `raw info 2`")
		}
	})

	t.Run("invalid_format", func(t *testing.T) {
		err := typego.NewErrorFromError(errors.New("error: code=01"))

		if errCode := err.GetCode(); errCode != "" {
			log.Fatal("`errCode` must be ``")
		}

		if errMessage := err.GetMessage(); errMessage != "" {
			log.Fatal("`errMessage` must be ``")
		}

		if errHttpStatus := err.GetHttpStatus(); errHttpStatus != 0 {
			log.Fatal("`errHttpStatus` must be `0`")
		}

		if errRPCStatus := err.GetRPCStatus(); errRPCStatus != 0 {
			log.Fatal("`errRPCStatus` must be `0`")
		}

		if errInfoLen := len(err.GetInfo()); errInfoLen != 0 {
			log.Fatal("`errInfoLen` must be `0`")
		}
	})
}

func TestErrorModel_AsGlobalVariable(t *testing.T) {
	errMap := make(map[string]string)
	err := typego.NewError("01", "")

	var wg sync.WaitGroup
	var lock = sync.RWMutex{}

	wg.Add(1)

	go func(w *sync.WaitGroup) {
		defer w.Done()

		e := err.ChangeCode("02")
		time.Sleep(2 * time.Second)

		lock.Lock()
		errMap["1"] = e.GetCode()
		lock.Unlock()
	}(&wg)

	wg.Add(1)

	go func(w *sync.WaitGroup) {
		defer w.Done()

		time.Sleep(1 * time.Second)
		e := err.ChangeCode("03")

		lock.Lock()
		errMap["2"] = e.GetCode()
		lock.Unlock()
	}(&wg)

	wg.Wait()

	if errCode := errMap["1"]; errCode != "02" {
		log.Fatal("`errCode` must be `02`")
	}

	if errCode2 := errMap["2"]; errCode2 != "03" {
		log.Fatal("`errCode2` must be `03`")
	}
}
