package capsule

import (
	"strconv"

	"github.com/valyala/fastjson"
)

//export hostHTTP
func hostHTTP(requestPosition, requestLength uint32, returnValuePosition **uint32, returnValueLength *uint32) uint32

// HTTP is an helper to use the hostHTTP function
func HTTP(request HTTPRequest) (HTTPResponse, error) {

	// Build the Json String from the request
	// (to be sent to the host)
	var body string
	var isJSON = false
	/*
	if request.Body != "" {
		body = request.Body
	} else if request.JSONBody != "" {
		isJSON = true
		body = request.JSONBody
	} else if request.TextBody != "" {
		body = request.TextBody
	} else {
		body=""
	}
	*/

	if request.JSONBody != "" {
		isJSON = true
		body = request.JSONBody
	} else if request.TextBody != "" {
		body = request.TextBody
	} else {
		body=""
	}	

	// TODO: we can only have one Body...
	var jsonHTTPRequest string
	/*
	if isJSON {
		jsonHTTPRequest = `{"Body":"","TextBody":"","JSONBody":` + body + `,"URI":"` + request.URI + `","Headers":` + request.Headers + `,"Method":"` + request.Method + `"}`
	} else {
		// add double quotes for body
		jsonHTTPRequest = `{"Body":"","JSONBody":{},"TextBody":"` + body + `","URI":"` + request.URI + `","Headers":` + request.Headers + `,"Method":"` + request.Method + `"}`
	}
	*/
	
	if isJSON {
		jsonHTTPRequest = `{"TextBody":"","JSONBody":` + body + `,"URI":"` + request.URI + `","Headers":` + request.Headers + `,"Method":"` + request.Method + `"}`
	} else {
		// add double quotes for body
		jsonHTTPRequest = `{"JSONBody":{},"TextBody":"` + body + `","URI":"` + request.URI + `","Headers":` + request.Headers + `,"Method":"` + request.Method + `"}`
	}
	
	jsonHTTPRequestPosition, jsonHTTPRequestSize := getBufferPosSize([]byte(jsonHTTPRequest))

	// This will be use to get the response from the host
	var responseBufferPtr *uint32
	var responseBufferSize uint32

	// Send the lessage to the host
	// (call the host function)
	hostHTTP(jsonHTTPRequestPosition, jsonHTTPRequestSize, &responseBufferPtr, &responseBufferSize)

	bufferResponseFromHost := readBufferFromMemory(responseBufferPtr, responseBufferSize)

	// Read in memory the value returned by the host
	JSONResponseBuffer, err := Result(bufferResponseFromHost)
	if err != nil {
		return HTTPResponse{}, err
	}

	//Print("✋ Response from the host (display by the guest): ")
	//Print(string(JSONResponseBuffer))

	// Parse the response
	parser := fastjson.Parser{}

	JSONData, err := parser.ParseBytes(JSONResponseBuffer)
	if err != nil {
		return HTTPResponse{}, err
	}

	// Extract data to construct the response
	jsonObject, err := JSONData.Object()
	if err != nil {
		//Log("cannot obtain object from json value: " + err.Error())
		return HTTPResponse{}, err
	}

	var JSONBody string
	var TextBody string
	//var Body string
	var Headers string
	var StatusCode string

	jsonObject.Visit(func(k []byte, v *fastjson.Value) {
		switch string(k) {
		case "JSONBody":
			JSONBody = v.String()
		case "TextBody":
			TextBody = v.String()
		//case "Body":
		//	Body = v.String()
		case "Headers":
			Headers = v.String()
		case "StatusCode":
			StatusCode = v.String()
		}
	})
	statusCode, _ := strconv.Atoi(StatusCode)

	return HTTPResponse{
		JSONBody:   JSONBody,
		TextBody:   TextBody,
		//Body:       Body,
		Headers:    Headers,
		StatusCode: statusCode,
	}, nil
}

