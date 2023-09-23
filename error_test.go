package typego_test

import (
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/dalikewara/typego"
	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	assert.NotNil(t, typego.NewError("", ""))
	assert.NotNil(t, typego.NewError("01", "general error"))
}

func TestErrorModel_ChangeCode(t *testing.T) {
	assert.NotNil(t, typego.NewError("", "").ChangeCode("01"))
}

func TestErrorModel_ChangeMessage(t *testing.T) {
	assert.NotNil(t, typego.NewError("", "").ChangeMessage("general error"))
}

func TestErrorModel_AddInfo(t *testing.T) {
	assert.NotNil(t, typego.NewError("", "").AddInfo(errors.New("raw error")))
}

func TestErrorModel_GetCode(t *testing.T) {
	assert.Equal(t, "01", typego.NewError("01", "").GetCode())
	assert.Equal(t, "02", typego.NewError("01", "").ChangeCode("02").GetCode())
}

func TestErrorModel_GetMessage(t *testing.T) {
	assert.Equal(t, "general error", typego.NewError("", "general error").GetMessage())
	assert.Equal(t, "general error 2", typego.NewError("", "general error").ChangeMessage("general error 2").GetMessage())
}

func TestErrorMessage_AddInfo(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		assert.Equal(t, []string{"raw error", "raw error 2", "raw error 3"}, typego.NewError("", "").AddInfo(errors.New("raw error"), errors.New("raw error 2")).AddInfo(errors.New("raw error 3")).GetInfo())
	})

	t.Run("string", func(t *testing.T) {
		assert.Equal(t, []string{"raw error", "raw error 2", "raw error 3"}, typego.NewError("", "").AddInfo("raw error", "raw error 2").AddInfo("raw error 3").GetInfo())
	})

	t.Run("any", func(t *testing.T) {
		assert.Equal(t, []string{"1"}, typego.NewError("", "").AddInfo(1).GetInfo())
	})
}

func TestErrorModel_Copy(t *testing.T) {
	t.Run("editing_global_variable_without_copying_the_object", func(t *testing.T) {
		// in this skenario, we assume that editing global variable without copying the object could cause such condition like race condition

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

		assert.Equal(t, "03", errMap["1"]) // false expected, should return 02
		assert.Equal(t, "03", errMap["2"]) // true expected, return 03
	})

	t.Run("editing_global_variable_by_copying_the_object", func(t *testing.T) {
		errMap := make(map[string]string)
		err := typego.NewError("01", "")

		var wg sync.WaitGroup
		var lock = sync.RWMutex{}

		wg.Add(1)
		go func(w *sync.WaitGroup) {
			defer w.Done()

			e := err.Copy().ChangeCode("02")
			time.Sleep(2 * time.Second)

			lock.Lock()
			errMap["1"] = e.GetCode()
			lock.Unlock()
		}(&wg)

		wg.Add(1)
		go func(w *sync.WaitGroup) {
			defer w.Done()

			time.Sleep(1 * time.Second)
			e := err.Copy().ChangeCode("03")

			lock.Lock()
			errMap["2"] = e.GetCode()
			lock.Unlock()
		}(&wg)

		wg.Wait()

		assert.Equal(t, "02", errMap["1"])
		assert.Equal(t, "03", errMap["2"])
	})
}

func TestErrorModel_Error(t *testing.T) {
	assert.Equal(t, "error: code=01, message=general error, info=[raw error]", typego.NewError("01", "general error").AddInfo(errors.New("raw error")).Error())
}
