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
	AddInfo(info ...interface{}) Error

	// AddDebug adds debug information and returns its instance
	AddDebug(debug ...interface{}) Error

	// SetHttpStatus sets error http status and returns its instance
	SetHttpStatus(httpStatus int) Error

	// SetRPCStatus sets error rpc status and returns its instance
	SetRPCStatus(rpcStatus int) Error

	// Log logs the error and return its instance
	Log() Error

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

	// Error returns error string
	Error() string
}

type errorModel struct {
	Code       string   `json:"code"`
	Message    string   `json:"message"`
	Info       []string `json:"info"`
	HttpStatus int      `json:"http_status,omitempty"`
	RPCStatus  int      `json:"rpc_status,omitempty"`
	Debug      []string `json:"debug,omitempty"`
}

func (e errorModel) ChangeCode(code string) Error {
	e.Code = code

	return e
}

func (e errorModel) ChangeMessage(message string) Error {
	e.Message = message

	return e
}

func (e errorModel) AddInfo(info ...interface{}) Error {
	for _, i := range info {
		if assertedString, ok := i.(string); ok {
			e.Info = append(e.Info, jsonStringCleaner(assertedString))

			continue
		}

		if assertedError, ok := i.(error); ok {
			e.Info = append(e.Info, jsonStringCleaner(assertedError.Error()))

			continue
		}

		e.Info = append(e.Info, fmt.Sprintf("%+v", i))
	}

	return e
}

func (e errorModel) AddDebug(debug ...interface{}) Error {
	for _, i := range debug {
		if assertedString, ok := i.(string); ok {
			e.Debug = append(e.Debug, jsonStringCleaner(assertedString))

			continue
		}

		if assertedError, ok := i.(error); ok {
			e.Debug = append(e.Debug, jsonStringCleaner(assertedError.Error()))

			continue
		}

		e.Debug = append(e.Debug, fmt.Sprintf("%+v", i))
	}

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

func (e errorModel) Log() Error {
	errorLogHandler(e)

	return e
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

func (e errorModel) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		return "invalid: " + err.Error()
	}

	return "error: " + string(b)
}

// NewError generates new typego.Error
func NewError(code string, message string) Error {
	return &errorModel{
		Code:    code,
		Message: message,
	}
}

// NewErrorFromError generates new typego.Error from an error. The error.Error() must have the same string format as
// typego.Error.Error(), otherwise, typego.Error will return incorrect value
func NewErrorFromError(err error) Error {
	var e errorModel

	errStr := err.Error()

	if len(errStr) > 7 {
		if er := json.Unmarshal([]byte(errStr[7:]), &e); er != nil {
			return &e
		}
	}

	return &e
}

// jsonStringCleaner cleans json string from double quotes (")
func jsonStringCleaner(jsonString string) string {
	var cleanedJSONString string

	length := len(jsonString)
	lengthPair := length - 1
	indexFlag := 4

	for i := 0; i < length; i++ {
		f := i + indexFlag

		if f > length {
			i -= 1
			indexFlag -= 1
			continue
		}

		var toBeCleaned string

		if i == lengthPair {
			toBeCleaned = jsonString[i:]
		} else {
			toBeCleaned = jsonString[i:f]
		}

		if toBeCleaned == "\"],\"" {
			cleanedJSONString += "], "
			i += indexFlag - 1
			indexFlag = 4

			continue
		} else if toBeCleaned == "\":[\"" {
			cleanedJSONString += ": ["
			i += indexFlag - 1
			indexFlag = 4

			continue
		} else if toBeCleaned == "\",\"" {
			cleanedJSONString += ", "
			i += indexFlag - 1
			indexFlag = 4

			continue
		} else if toBeCleaned == "\":\"" {
			cleanedJSONString += ": "
			i += indexFlag - 1
			indexFlag = 4

			continue
		} else if toBeCleaned == "{\"" {
			cleanedJSONString += "{"
			i += indexFlag - 1
			indexFlag = 4

			continue
		} else if toBeCleaned == "\":" {
			cleanedJSONString += ": "
			i += indexFlag - 1
			indexFlag = 4

			continue
		} else if toBeCleaned == ",\"" {
			cleanedJSONString += ", "
			i += indexFlag - 1
			indexFlag = 4

			continue
		}

		if indexFlag == 1 {
			cleanedJSONString += toBeCleaned
			i += indexFlag - 1
			indexFlag = 4

			continue
		}

		indexFlag -= 1

		if indexFlag < 1 {
			indexFlag = 4

			continue
		}

		i -= 1
	}

	return cleanedJSONString
}
