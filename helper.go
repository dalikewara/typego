package typego

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
