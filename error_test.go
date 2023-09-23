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
	assert.NotNil(t, typego.NewError())
}

func TestErrorModel_SetCode(t *testing.T) {
	assert.NotNil(t, typego.NewError().SetCode("01"))
}

func TestErrorModel_SetMessage(t *testing.T) {
	assert.NotNil(t, typego.NewError().SetMessage("general error"))
}

func TestErrorModel_SetInfo(t *testing.T) {
	assert.NotNil(t, typego.NewError().SetInfo(errors.New("raw error")))
}

func TestErrorModel_GetCode(t *testing.T) {
	assert.Equal(t, "01", typego.NewError().SetCode("01").GetCode())
}

func TestErrorModel_GetMessage(t *testing.T) {
	assert.Equal(t, "general error", typego.NewError().SetMessage("general error").GetMessage())
}

func TestErrorMessage_GetInfo(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		assert.Equal(t, []string{"raw error", "raw error 2", "raw error 3"}, typego.NewError().SetInfo(errors.New("raw error"), errors.New("raw error 2")).SetInfo(errors.New("raw error 3")).GetInfo())
	})

	t.Run("string", func(t *testing.T) {
		assert.Equal(t, []string{"raw error", "raw error 2", "raw error 3"}, typego.NewError().SetInfo("raw error", "raw error 2").SetInfo("raw error 3").GetInfo())
	})

	t.Run("any", func(t *testing.T) {
		assert.Equal(t, []string{"1"}, typego.NewError().SetInfo(1).GetInfo())
	})
}

func TestErrorModel_Copy(t *testing.T) {
	t.Run("editing_global_variable_without_copying_the_object", func(t *testing.T) {
		// in this skenario, we assume that editing global variable without copying the object could cause such condition like race condition

		errMap := make(map[string]string)
		err := typego.NewError().SetCode("01")

		var wg sync.WaitGroup
		var lock = sync.RWMutex{}

		wg.Add(1)
		go func(w *sync.WaitGroup) {
			defer w.Done()

			e := err.SetCode("02")
			time.Sleep(2 * time.Second)

			lock.Lock()
			errMap["1"] = e.GetCode()
			lock.Unlock()
		}(&wg)

		wg.Add(1)
		go func(w *sync.WaitGroup) {
			defer w.Done()

			time.Sleep(1 * time.Second)
			e := err.SetCode("03")

			lock.Lock()
			errMap["2"] = e.GetCode()
			lock.Unlock()
		}(&wg)

		wg.Wait()

		assert.Equal(t, "03", errMap["1"]) // false expected, should has 02
		assert.Equal(t, "03", errMap["2"]) // true expected, has 03
	})

	t.Run("editing_global_variable_by_copying_the_object", func(t *testing.T) {
		errMap := make(map[string]string)
		err := typego.NewError().SetCode("01")

		var wg sync.WaitGroup
		var lock = sync.RWMutex{}

		wg.Add(1)
		go func(w *sync.WaitGroup) {
			defer w.Done()

			e := err.Copy().SetCode("02")
			time.Sleep(2 * time.Second)

			lock.Lock()
			errMap["1"] = e.GetCode()
			lock.Unlock()
		}(&wg)

		wg.Add(1)
		go func(w *sync.WaitGroup) {
			defer w.Done()

			time.Sleep(1 * time.Second)
			e := err.Copy().SetCode("03")

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
	assert.Equal(t, "error: code=01, message=general error", typego.NewError().SetCode("01").SetMessage("general error").Error())
}
