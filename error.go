package typego

import (
	"encoding/json"
	"fmt"
)

type Error interface {
	// ChangeCode changes error code and returns its instance
	ChangeCode(code string) Error

	// ChangeMessage changes error message and returns its instance
	ChangeMessage(message string) Error

	// AddInfo adds error information and returns its instance
	AddInfo(info ...any) Error

	// AddDebug adds debug information and returns its instance
	AddDebug(debug ...any) Error

	// SetProcessID sets process id
	SetProcessID(processID string) Error

	// SetProcessName sets process name
	SetProcessName(processName string) Error

	// SetHttpStatus sets error http status and returns its instance
	SetHttpStatus(httpStatus int) Error

	// SetRPCStatus sets error rpc status and returns its instance
	SetRPCStatus(rpcStatus int) Error

	// GetProcessID gets process id
	GetProcessID() string

	// GetProcessName gets process name
	GetProcessName() string

	// GetCode gets error code
	GetCode() string

	// GetMessage gets error message
	GetMessage() string

	// GetInfo gets error information
	GetInfo() []string

	// GetDebug gets debug information
	GetDebug() []string

	// GetHttpStatus gets error http status
	GetHttpStatus() int

	// GetRPCStatus gets error rpc status
	GetRPCStatus() int

	// Log logs the error and return its instance
	Log() Error

	// Error returns error string
	Error() string
}

type errorModel struct {
	Level       string   `json:"level"`
	ProcessID   string   `json:"process_id,omitempty"`
	ProcessName string   `json:"process_name,omitempty"`
	Code        string   `json:"code"`
	Message     string   `json:"message"`
	Info        []string `json:"info"`
	HttpStatus  int      `json:"http_status,omitempty"`
	RPCStatus   int      `json:"rpc_status,omitempty"`
	Debug       []string `json:"debug,omitempty"`
}

func (e errorModel) SetProcessID(processID string) Error {
	e.ProcessID = processID
	return e
}

func (e errorModel) ChangeCode(code string) Error {
	e.Code = code
	return e
}

func (e errorModel) ChangeMessage(message string) Error {
	e.Message = message
	return e
}

func (e errorModel) AddInfo(info ...any) Error {
	additionalInfo := make([]string, 0, len(info))

	for _, i := range info {
		switch v := i.(type) {
		case string:
			additionalInfo = append(additionalInfo, JSONStringCleaner(v))
		case error:
			additionalInfo = append(additionalInfo, JSONStringCleaner(v.Error()))
		default:
			jsonValue, err := json.Marshal(v)
			if err != nil {
				additionalInfo = append(additionalInfo, fmt.Sprintf("%+v", v))
			} else {
				additionalInfo = append(additionalInfo, string(jsonValue))
			}
		}
	}

	e.Info = append(e.Info, additionalInfo...)

	return e
}

func (e errorModel) AddDebug(debug ...any) Error {
	additionalDebug := make([]string, 0, len(debug))

	for _, i := range debug {
		switch v := i.(type) {
		case string:
			additionalDebug = append(additionalDebug, JSONStringCleaner(v))
		case error:
			additionalDebug = append(additionalDebug, JSONStringCleaner(v.Error()))
		default:
			jsonValue, err := json.Marshal(v)
			if err != nil {
				additionalDebug = append(additionalDebug, fmt.Sprintf("%+v", v))
			} else {
				additionalDebug = append(additionalDebug, string(jsonValue))
			}
		}
	}

	e.Debug = append(e.Debug, additionalDebug...)

	return e
}

func (e errorModel) SetProcessName(processName string) Error {
	e.ProcessName = processName
	return e
}

func (e errorModel) SetHttpStatus(httpStatus int) Error {
	e.HttpStatus = httpStatus
	return e
}

func (e errorModel) SetRPCStatus(rpcStatus int) Error {
	e.RPCStatus = rpcStatus
	return e
}

func (e errorModel) GetProcessID() string {
	return e.ProcessID
}

func (e errorModel) GetProcessName() string {
	return e.ProcessName
}

func (e errorModel) GetCode() string {
	return e.Code
}

func (e errorModel) GetMessage() string {
	return e.Message
}

func (e errorModel) GetInfo() []string {
	return e.Info
}

func (e errorModel) GetDebug() []string {
	return e.Debug
}

func (e errorModel) GetHttpStatus() int {
	return e.HttpStatus
}

func (e errorModel) GetRPCStatus() int {
	return e.RPCStatus
}

func (e errorModel) Log() Error {
	errorLogHandler(e)
	return e
}

func (e errorModel) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		return err.Error()
	}

	return string(b)
}

// NewError generates new typego.Error
func NewError(code string, message string) Error {
	return &errorModel{
		Level:   "error",
		Code:    code,
		Message: message,
	}
}

// NewErrorFromError generates new typego.Error from an error. The error.Error() must have the same string format as
// typego.Error.Error(), otherwise, typego.Error will return incorrect value
func NewErrorFromError(err error) Error {
	var e errorModel

	if er := json.Unmarshal([]byte(err.Error()), &e); er != nil {
		return &e
	}

	return &e
}
