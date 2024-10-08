package typego_test

import (
	"errors"
	"fmt"
	"github.com/dalikewara/typego"
	"log"
	"testing"
)

func TestNewInfo(t *testing.T) {
	if info := typego.NewInfo(); info == nil {
		log.Fatal("`info` must not nil")
	}
}

func TestInfoModel_AddInfo(t *testing.T) {
	if info := typego.NewInfo().AddInfo(errors.New("raw info")); info == nil {
		log.Fatal("`info` must not nil")
	}
}

func TestInfoModel_AddDebug(t *testing.T) {
	if info := typego.NewInfo().AddDebug(errors.New("raw info")); info == nil {
		log.Fatal("`info` must not nil")
	}
}

func TestInfoModel_SetProcessID(t *testing.T) {
	if info := typego.NewInfo().SetProcessID("123"); info == nil {
		log.Fatal("`info` must not nil")
	}
}

func TestInfoModel_SetProcessName(t *testing.T) {
	if info := typego.NewInfo().SetProcessName("test"); info == nil {
		log.Fatal("`info` must not nil")
	}
}

func TestInfoModel_GetProcessID(t *testing.T) {
	if infoProcessID := typego.NewInfo().SetProcessID("123").GetProcessID(); infoProcessID != "123" {
		log.Fatal("`infoProcessID` must be `123`")
	}
}

func TestInfoModel_GetProcessName(t *testing.T) {
	if infoProcessName := typego.NewInfo().SetProcessName("test").GetProcessName(); infoProcessName != "test" {
		log.Fatal("`infoProcessName` must be `test`")
	}
}

func TestInfoModel_GetInfo(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		if infoInfos := typego.NewInfo().AddInfo(errors.New("raw info"), errors.New("raw info 2")).AddInfo(errors.New("raw info 3")).GetInfo(); fmt.Sprintf("%v", infoInfos) != fmt.Sprintf("%v", []string{"raw info", "raw info 2", "raw info 3"}) {
			log.Fatal("`infoInfos` must be `[]string{\"raw info\", \"raw info 2\", \"raw info 3\"}`")
		}
	})

	t.Run("string", func(t *testing.T) {
		if infoInfos := typego.NewInfo().AddInfo("raw info", "raw info 2").AddInfo("raw info 3").GetInfo(); fmt.Sprintf("%v", infoInfos) != fmt.Sprintf("%v", []string{"raw info", "raw info 2", "raw info 3"}) {
			log.Fatal("`infoInfos` must be `[]string{\"raw info\", \"raw info 2\", \"raw info 3\"}`")
		}
	})

	t.Run("any", func(t *testing.T) {
		if infoInfos := typego.NewInfo().AddInfo(1).GetInfo(); fmt.Sprintf("%v", infoInfos) != fmt.Sprintf("%v", []string{"1"}) {
			log.Fatal("`infoInfos` must be `[]string{\"1\"}`")
		}
	})
}

func TestInfoModel_GetDebug(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		if infoDebugs := typego.NewInfo().AddDebug(errors.New("raw info"), errors.New("raw info 2")).AddDebug(errors.New("raw info 3")).GetDebug(); fmt.Sprintf("%v", infoDebugs) != fmt.Sprintf("%v", []string{"raw info", "raw info 2", "raw info 3"}) {
			log.Fatal("`infoDebugs` must be `[]string{\"raw info\", \"raw info 2\", \"raw info 3\"}`")
		}
	})

	t.Run("string", func(t *testing.T) {
		if infoDebugs := typego.NewInfo().AddDebug("raw info", "raw info 2").AddDebug("raw info 3").GetDebug(); fmt.Sprintf("%v", infoDebugs) != fmt.Sprintf("%v", []string{"raw info", "raw info 2", "raw info 3"}) {
			log.Fatal("`infoDebugs` must be `[]string{\"raw info\", \"raw info 2\", \"raw info 3\"}`")
		}
	})

	t.Run("any", func(t *testing.T) {
		if infoDebugs := typego.NewInfo().AddDebug(1).GetDebug(); fmt.Sprintf("%v", infoDebugs) != fmt.Sprintf("%v", []string{"1"}) {
			log.Fatal("`infoDebugs` must be `[]string{\"1\"}`")
		}
	})
}

func TestInfoModel_Log(t *testing.T) {
	_ = typego.NewInfo().Log()
}

func TestInfoModel_String(t *testing.T) {
	if info := typego.NewInfo().String(); info != "{\"level\":\"info\",\"info\":null}" {
		log.Fatal("`info` must be `{\"level\":\"info\",\"info\":null}`")
	}
}
