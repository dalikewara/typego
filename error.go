package typego

import "fmt"

type Error interface {
	SetCode(code string) Error
	SetMessage(message string) Error
	GetCode() string
	GetMessage() string
	Copy() Error
	Error() string
}

type errorModel struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// SetCode sets error code
func (e *errorModel) SetCode(code string) Error {
	e.Code = code
	return e
}

// SetMessage sets error message
func (e *errorModel) SetMessage(message string) Error {
	e.Message = message
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

// Copy copies the error object and returns the new one
func (e *errorModel) Copy() Error {
	copied := *e
	return &copied
}

// Error returns error string
func (e *errorModel) Error() string {
	return fmt.Sprintf("error: code=%s, message=%s", e.Code, e.Message)
}

// NewError generates new typego.Error
func NewError() Error {
	return &errorModel{}
}
