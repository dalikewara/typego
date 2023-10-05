package typego

import (
	"encoding/json"
	"fmt"
)

type Error interface {
	ChangeCode(code string) Error
	ChangeMessage(message string) Error
	AddInfo(info ...interface{}) Error
	SetHttpStatus(httpStatus int) Error
	SetRPCStatus(rpcStatus int) Error
	GetCode() string
	GetMessage() string
	GetInfo() []string
	GetHttpStatus() int
	GetRPCStatus() int
	Copy() Error
	Error() string
}

type errorModel struct {
	Code       string   `json:"code"`
	Message    string   `json:"message"`
	Info       []string `json:"info"`
	HttpStatus int      `json:"http_status"`
	RPCStatus  int      `json:"rpc_status"`
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

// AddInfo adds error information
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

// SetHttpStatus sets error http status
func (e *errorModel) SetHttpStatus(httpStatus int) Error {
	e.HttpStatus = httpStatus
	return e
}

// SetRPCStatus sets error rpc status
func (e *errorModel) SetRPCStatus(rpcStatus int) Error {
	e.RPCStatus = rpcStatus
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

// GetHttpStatus gets error http status
func (e *errorModel) GetHttpStatus() int {
	return e.HttpStatus
}

// GetRPCStatus gets error rpc status
func (e *errorModel) GetRPCStatus() int {
	return e.RPCStatus
}

// Copy copies the error object and returns the new one
func (e *errorModel) Copy() Error {
	copied := *e
	return &copied
}

// Error returns error string
func (e *errorModel) Error() string {
	b, _ := json.Marshal(e)
	return "error: " + string(b)
}

// NewError generates new typego.Error
func NewError(code string, message string) Error {
	return &errorModel{
		Code:       code,
		Message:    message,
		HttpStatus: 500,
		RPCStatus:  13,
	}
}

// NewErrorFromError generates new typego.Error from an error. The error.Error() must has the same string format as typego.Error.Error(), otherwise, typego.Error will return incorrect value
func NewErrorFromError(err error) Error {
	var e errorModel

	errStr := err.Error()

	if len(errStr) > 7 {
		json.Unmarshal([]byte(errStr[7:]), &e)
	}

	if e.HttpStatus == 0 {
		e.HttpStatus = 500
	}

	if e.RPCStatus == 0 {
		e.RPCStatus = 13
	}

	return &e
}
