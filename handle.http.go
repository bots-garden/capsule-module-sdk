package capsule

import (
	"strconv"

	"github.com/valyala/fastjson"
)

// HTTPRequest is the data of the http request
type HTTPRequest struct {
	Body string
	URI      string
	Method   string
	Headers  string
}
//JsonHTTPRequest ?

// HTTPResponse is the data of the http response
type HTTPResponse struct {
	Body string
	Headers  string
	StatusCode int
}

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
		URI:      string(JSONData.GetStringBytes("URI")),
		Method:   string(JSONData.GetStringBytes("Method")),
		Headers:  string(JSONData.GetStringBytes("Headers")),
	}

	// call the handle function
	retValue, err := handleHTTPFunction(httpRequestParam)

	//TODO: return an HTTPResponse
	// return failure or success
	if err != nil {
		return failure([]byte(err.Error()))
	}

	jsonHTTPResponse := `{"Body":"`+retValue.Body+`","Headers":`+retValue.Headers+`,"StatusCode":`+strconv.Itoa(retValue.StatusCode)+`}`

	// first byte == 82
	return success([]byte(jsonHTTPResponse))

}