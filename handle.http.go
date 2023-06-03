package capsule

import (
	"encoding/base64"
	"strconv"

	"github.com/valyala/fastjson"
)



var handleHTTPFunction func(param HTTPRequest) (HTTPResponse, error)

// SetHandleHTTP sets the handle function
func SetHandleHTTP(function func(param HTTPRequest) (HTTPResponse, error)) {
	handleHTTPFunction = function
}

//export callHandleHTTP
func callHandleHTTP(JSONDataPos *uint32, JSONDataSize uint32) uint64 {
	
	parser := fastjson.Parser{}

	JSONDataBuffer := readBufferFromMemory(JSONDataPos, JSONDataSize)
	JSONData, err := parser.ParseBytes(JSONDataBuffer)
	if err != nil {
		return failure([]byte(err.Error()))
	}
	httpRequestParam := HTTPRequest{
		Body: string(JSONData.GetStringBytes("Body")),
		//JSONBody: string(JSONData.GetStringBytes("JSONBody")), //! to use in the future
		//TextBody: string(JSONData.GetStringBytes("TextBody")), //! to use in the future
		URI:      string(JSONData.GetStringBytes("URI")),
		Method:   string(JSONData.GetStringBytes("Method")),
		Headers:  string(JSONData.GetStringBytes("Headers")),
	}

	// call the handle function
	retValue, err := handleHTTPFunction(httpRequestParam)

	// return failure or success
	if err != nil {
		return failure([]byte(err.Error()))
	}

	var jsonBody string
	if len(retValue.JSONBody) == 0 {
		jsonBody = "{}"
	} else {
		jsonBody = retValue.JSONBody
	}
	
	var textBody string
	if len(retValue.TextBody) == 0 {
		textBody = ""
	} else {
		// avoid special characters in jsonString
		textBody = base64.StdEncoding.EncodeToString([]byte(retValue.TextBody))
		//textBody = retValue.TextBody
	}

	jsonHTTPResponse := `{"JSONBody":`+jsonBody+`,"TextBody":"`+textBody+`","Headers":`+retValue.Headers+`,"StatusCode":`+strconv.Itoa(retValue.StatusCode)+`}`

	// first byte == 82
	return success([]byte(jsonHTTPResponse))

}