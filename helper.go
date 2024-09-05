package typego

import "strings"

// jsonStringCleaner cleans json string from double quotes (")
func jsonStringCleaner(jsonString string) string {
	var builder strings.Builder

	length := len(jsonString)
	i := 0

	for i < length {
		if i+4 <= length {
			switch jsonString[i : i+4] {
			case "\"],\"":
				builder.WriteString("], ")
				i += 4
				continue
			case "\":[\"":
				builder.WriteString(": [")
				i += 4
				continue
			case "\",\"":
				builder.WriteString(", ")
				i += 4
				continue
			case "\":\"":
				builder.WriteString(": ")
				i += 4
				continue
			}
		}

		if i+3 <= length {
			switch jsonString[i : i+3] {
			case "{\"\"":
				builder.WriteString("{")
				i += 3
				continue
			case ",\"":
				builder.WriteString(", ")
				i += 2
				continue
			}
		}

		builder.WriteByte(jsonString[i])
		i++
	}

	return builder.String()
}
