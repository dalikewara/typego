package typego

import "fmt"

type Error interface {
	SetCode(code string) Error
	SetMessage(message string) Error
	SetInfo(info ...interface{}) Error
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

// SetInfo sets error info
func (e *errorModel) SetInfo(info ...interface{}) Error {
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
	return fmt.Sprintf("error: code=%s, message=%s", e.Code, e.Message)
}

// NewError generates new typego.Error
func NewError() Error {
	return &errorModel{}
}
