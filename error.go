package typego

import (
	"fmt"
	"strings"
)

type Error interface {
	ChangeCode(code string) Error
	ChangeMessage(message string) Error
	AddInfo(info ...interface{}) Error
	GetCode() string
	GetMessage() string
	GetInfo() []string
	Copy() Error
	Error() string
}

type errorModel struct {
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Info    []string `json:"info"`
}

// ChangeCode changes error code
func (e *errorModel) ChangeCode(code string) Error {
	e.Code = code
	return e
}

// ChangeMessage changes error message
func (e *errorModel) ChangeMessage(message string) Error {
	e.Message = message
	return e
}

// AddInfo adds error info
func (e *errorModel) AddInfo(info ...interface{}) Error {
	for _, i := range info {
		if asserted, ok := i.(error); ok {
			e.Info = append(e.Info, asserted.Error())
		} else if asserted, ok := i.(string); ok {
			e.Info = append(e.Info, asserted)
		} else {
			e.Info = append(e.Info, fmt.Sprintf("%+v", i))
		}
	}
	return e
}

// GetCode gets error code
func (e *errorModel) GetCode() string {
	return e.Code
}

// GetMessage gets error message
func (e *errorModel) GetMessage() string {
	return e.Message
}

// GetInfo gets error information
func (e *errorModel) GetInfo() []string {
	return e.Info
}

// Copy copies the error object and returns the new one
func (e *errorModel) Copy() Error {
	copied := *e
	return &copied
}

// Error returns error string
func (e *errorModel) Error() string {
	return fmt.Sprintf("error: code=%s, message=%s, info=[%s]", e.Code, e.Message, strings.Join(e.Info, ", "))
}

// NewError generates new typego.Error
func NewError(code string, message string) Error {
	return &errorModel{
		Code:    code,
		Message: message,
	}
}
