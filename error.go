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
	Error() string
	Log()
}

type errorModel struct {
	Code       string   `json:"code"`
	Message    string   `json:"message"`
	Info       []string `json:"info"`
	HttpStatus int      `json:"http_status,omitempty"`
	RPCStatus  int      `json:"rpc_status,omitempty"`
}

// ChangeCode changes error code
func (e errorModel) ChangeCode(code string) Error {
	e.Code = code
	return e
}

// ChangeMessage changes error message
func (e errorModel) ChangeMessage(message string) Error {
	e.Message = message
	return e
}

// AddInfo adds error information
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

// SetHttpStatus sets error http status
func (e errorModel) SetHttpStatus(httpStatus int) Error {
	e.HttpStatus = httpStatus
	return e
}

// SetRPCStatus sets error rpc status
func (e errorModel) SetRPCStatus(rpcStatus int) Error {
	e.RPCStatus = rpcStatus
	return e
}

// GetCode gets error code
func (e errorModel) GetCode() string {
	return e.Code
}

// GetMessage gets error message
func (e errorModel) GetMessage() string {
	return e.Message
}

// GetInfo gets error information
func (e errorModel) GetInfo() []string {
	return e.Info
}

// GetHttpStatus gets error http status
func (e errorModel) GetHttpStatus() int {
	return e.HttpStatus
}

// GetRPCStatus gets error rpc status
func (e errorModel) GetRPCStatus() int {
	return e.RPCStatus
}

// Error returns error string
func (e errorModel) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		return "invalid: " + err.Error()
	}

	return "error: " + string(b)
}

func (e errorModel) Log() {
	fmt.Println(fmt.Sprintf("%+v", e))
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
