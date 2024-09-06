package typego

import (
	"encoding/json"
	"fmt"
)

type Info interface {
	// AddInfo adds information and returns its instance
	AddInfo(info ...interface{}) Info

	// AddDebug adds information debug and returns its instance
	AddDebug(debug ...interface{}) Info

	// SetProcessID sets process id
	SetProcessID(processID string) Info

	// SetProcessName sets process name
	SetProcessName(processName string) Info

	// GetProcessID gets process id
	GetProcessID() string

	// GetProcessName gets process name
	GetProcessName() string

	// GetInfo gets information
	GetInfo() []string

	// GetDebug gets information debug
	GetDebug() []string

	// Log logs the information and return its instance
	Log() Info

	// String returns the information in string
	String() string
}

type infoModel struct {
	Level       string   `json:"level"`
	ProcessID   string   `json:"process_id,omitempty"`
	ProcessName string   `json:"process_name,omitempty"`
	Info        []string `json:"info"`
	Debug       []string `json:"debug,omitempty"`
}

func (i infoModel) AddInfo(info ...interface{}) Info {
	additionalInfo := make([]string, 0, len(info))

	for _, j := range info {
		switch v := j.(type) {
		case string:
			additionalInfo = append(additionalInfo, JSONStringCleaner(v))
		case error:
			additionalInfo = append(additionalInfo, JSONStringCleaner(v.Error()))
		default:
			jsonValue, err := json.Marshal(v)
			if err != nil {
				additionalInfo = append(additionalInfo, fmt.Sprintf("%+v", v))
			} else {
				additionalInfo = append(additionalInfo, JSONStringCleaner(string(jsonValue)))
			}
		}
	}

	i.Info = append(i.Info, additionalInfo...)

	return i
}

func (i infoModel) AddDebug(debug ...interface{}) Info {
	additionalDebug := make([]string, 0, len(debug))

	for _, j := range debug {
		switch v := j.(type) {
		case string:
			additionalDebug = append(additionalDebug, JSONStringCleaner(v))
		case error:
			additionalDebug = append(additionalDebug, JSONStringCleaner(v.Error()))
		default:
			jsonValue, err := json.Marshal(v)
			if err != nil {
				additionalDebug = append(additionalDebug, fmt.Sprintf("%+v", v))
			} else {
				additionalDebug = append(additionalDebug, JSONStringCleaner(string(jsonValue)))
			}
		}
	}

	i.Debug = append(i.Debug, additionalDebug...)

	return i
}

func (i infoModel) SetProcessID(processID string) Info {
	i.ProcessID = processID
	return i
}

func (i infoModel) SetProcessName(processName string) Info {
	i.ProcessName = processName
	return i
}

func (i infoModel) GetProcessID() string {
	return i.ProcessID
}

func (i infoModel) GetProcessName() string {
	return i.ProcessName
}

func (i infoModel) GetInfo() []string {
	return i.Info
}

func (i infoModel) GetDebug() []string {
	return i.Debug
}

func (i infoModel) Log() Info {
	infoLogHandler(i)
	return i
}

func (i infoModel) String() string {
	b, err := json.Marshal(i)
	if err != nil {
		return err.Error()
	}

	return string(b)
}

// NewInfo generates new typego.Info
func NewInfo() Info {
	return &infoModel{
		Level: "info",
	}
}
