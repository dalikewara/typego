package typego

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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
	httpStatus int
	rpcStatus  int
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
	e.httpStatus = httpStatus
	return e
}

// SetRPCStatus sets error rpc status
func (e *errorModel) SetRPCStatus(rpcStatus int) Error {
	e.rpcStatus = rpcStatus
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
	return e.httpStatus
}

// GetRPCStatus gets error rpc status
func (e *errorModel) GetRPCStatus() int {
	return e.rpcStatus
}

// Copy copies the error object and returns the new one
func (e *errorModel) Copy() Error {
	copied := *e
	return &copied
}

// Error returns error string
func (e *errorModel) Error() string {
	return fmt.Sprintf("error: code=%s, message=%s, httpStatus=%v, rpcStatus=%v, info=%s", e.Code, e.Message, e.httpStatus, e.rpcStatus, strings.Join(e.Info, ", info="))
}

// NewError generates new typego.Error
func NewError(code string, message string) Error {
	return &errorModel{
		Code:       code,
		Message:    message,
		httpStatus: 500,
		rpcStatus:  13,
	}
}

// NewErrorFromError generates new typego.Error from an error. The error.Error() must has the same string format as typego.Error.Error(), otherwise, typego.Error will return incorrect value. The expected string format is: `error: code=%s, message=%s, httpStatus=%v, rpcStatus=%v, info=%s`
func NewErrorFromError(err error) (Error, error) {
	var value string
	var flagCode, flagMessage, flagHttpStatus, flagRPCStatus int

	e := &errorModel{}
	errString := err.Error()
	length := len(errString)

	for i := 0; i <= length-1; i++ {
		if (i+11) <= (length-1) && errString[i:i+12] == "error: code=" {
			if flagCode != 0 {
				return nil, errors.New("the parameters value cannot contain this string: `error: code=`")
			}
			value = ""
			i += 11
			flagCode++
		} else if (i+9) <= (length-1) && errString[i:i+10] == ", message=" {
			if flagMessage != 0 {
				return nil, errors.New("the parameters value cannot contain this string: `, message=`")
			}
			e.ChangeCode(value)
			value = ""
			i += 9
			flagMessage++
		} else if (i+12) <= (length-1) && errString[i:i+13] == ", httpStatus=" {
			if flagHttpStatus != 0 {
				return nil, errors.New("the parameters value cannot contain this string: `, httpStatus=`")
			}
			e.ChangeMessage(value)
			value = ""
			i += 12
			flagHttpStatus++

		} else if (i+11) <= (length-1) && errString[i:i+12] == ", rpcStatus=" {
			if flagRPCStatus != 0 {
				return nil, errors.New("the parameters value cannot contain this string: `, rpcStatus=`")
			}
			n, _ := strconv.Atoi(value)
			if n == 0 {
				e.SetHttpStatus(500)
			} else {
				e.SetHttpStatus((int(n)))
			}
			value = ""
			i += 11
			flagRPCStatus++
		} else if (i+6) <= (length-1) && errString[i:i+7] == ", info=" {
			if e.GetRPCStatus() == 0 {
				n, _ := strconv.Atoi(value)
				if n == 0 {
					e.SetRPCStatus(13)
				} else {
					e.SetRPCStatus(int(n))
				}
			} else {
				e.AddInfo(value)
			}
			value = ""
			i += 6

		} else {
			value += string(errString[i])
			if i == length-1 {
				e.AddInfo(value)
				value = ""
			}
		}
	}

	return e, nil
}
