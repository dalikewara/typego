package typego

import "fmt"

type Info interface {
	// AddInfo adds information and returns its instance
	AddInfo(info ...interface{}) Info

	// AddDebug adds debug information and returns its instance
	AddDebug(debug ...interface{}) Info

	// SetProcessName sets process name
	SetProcessName(processName string) Info

	// Log logs the information and return its instance
	Log() Info

	// GetInfo gets information
	GetInfo() []string

	// GetDebug gets debug information
	GetDebug() []string
}

type infoModel struct {
	Level       string   `json:"level"`
	ProcessName string   `json:"process_name,omitempty"`
	Info        []string `json:"info"`
	Debug       []string `json:"debug,omitempty"`
}

func (i infoModel) AddInfo(info ...interface{}) Info {
	for _, j := range info {
		if assertedString, ok := j.(string); ok {
			i.Info = append(i.Info, jsonStringCleaner(assertedString))

			continue
		}

		if assertedError, ok := j.(error); ok {
			i.Info = append(i.Info, jsonStringCleaner(assertedError.Error()))

			continue
		}

		i.Info = append(i.Info, fmt.Sprintf("%+v", j))
	}

	return i
}

func (i infoModel) AddDebug(debug ...interface{}) Info {
	for _, j := range debug {
		if assertedString, ok := j.(string); ok {
			i.Debug = append(i.Debug, jsonStringCleaner(assertedString))

			continue
		}

		if assertedError, ok := j.(error); ok {
			i.Debug = append(i.Debug, jsonStringCleaner(assertedError.Error()))

			continue
		}

		i.Debug = append(i.Debug, fmt.Sprintf("%+v", j))
	}

	return i
}

func (i infoModel) SetProcessName(processName string) Info {
	i.ProcessName = processName

	return i
}

func (i infoModel) Log() Info {
	infoLogHandler(i)

	return i
}

func (i infoModel) GetInfo() []string {
	return i.Info
}

func (i infoModel) GetDebug() []string {
	return i.Debug
}

// NewInfo generates new typego.Info
func NewInfo() Info {
	return &infoModel{
		Level: "info",
	}
}
