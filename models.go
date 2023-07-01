package capsule

import "strings"

// HTTPRequest is the data of the http request
type HTTPRequest struct {
	Body    string
	JSONBody   string
	TextBody   string
	URI     string
	Method  string
	Headers string
}

// HTTPResponse is the data of the http response
type HTTPResponse struct {
	//Body    string
	JSONBody   string
	TextBody   string
	Headers    string
	StatusCode int
}

// strSeparator is use to create a string from a slice
// ex: "Content-Type":"application/json; charset=utf-8","PRIVATE-TOKEN":"glpat-mmysLDdMx532zxHgsgzF"
const strSeparator = ","
// fieldSeparator is used to create a slice from a map
// ex: "PRIVATE-TOKEN":"glpat-mmysLDdMx532zxHgsgzF"
const fieldSeparator = ":"
// quote is used to create a slice from a map
// to add double quote before and after a fieldname or a value
const quote = "\""


// createStringFromSlice creates a string from a slice of strings
func createStringFromSlice(strSlice []string, separator string) string {
	return strings.Join(strSlice[:], separator)
}

// createSliceFromMap creates a slice of strings from a maps of strings
func createSliceFromMap(strMap map[string]string) []string {
	var strSlice []string
	for field, value := range strMap {
		strSlice = append(strSlice, quote+field+quote+fieldSeparator+quote+value+quote)
	}
	return strSlice
}

// createSliceFromString creates a slice of string from a string
func createSliceFromString(str string, separator string) []string {
	return strings.Split(str, separator)
}
// createMapFromSlice creates a map from a slice of string
func createMapFromSlice(strSlice []string, separator string, remove string) map[string]string {
	strMap := make(map[string]string)
	for _, item := range strSlice {
		
		item = strings.ReplaceAll(item, remove, "")
		res := strings.Split(item, separator)

		if len(res) >= 2 { // for example: "https://gitlab.com" and separator is ":"
			strMap[res[0]] = strings.Join(res[1:], separator)
		} else {
			if len(res) > 1 {
				strMap[res[0]] = res[1]
			}
		}
	}
	return strMap
}

// GetHeaders return a map of headers from a string of headers
// transform a JSON string to a map[string]string
func GetHeaders(headers string) map[string]string {
	return createMapFromSlice(createSliceFromString(headers, ","), ":", `"`)
}

// SetHeaders return a string of headers from a map of headers
// transform a map[string]string to a JSON string
func SetHeaders(headers map[string]string) string {
	return "{"+createStringFromSlice(createSliceFromMap(headers), strSeparator)+"}"
}

