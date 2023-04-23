package capsule

import "github.com/valyala/fastjson"


var handleJSONFunction func(param *fastjson.Value) ([]byte, error)

// SetHandleJSON sets the handle function
func SetHandleJSON(function func(param *fastjson.Value) ([]byte, error)) {
	handleJSONFunction = function
}

//export callHandleJSON
func callHandleJSON(subjectPosition *uint32, length int) uint64 {
	// read the memory to get the parameter
	subjectBytes := read(subjectPosition, length)

	parser := fastjson.Parser{}
    jsonValue, err := parser.ParseBytes(subjectBytes)
	//TODO: handle error

	// call the handle function
	retValue, err := handleJSONFunction(jsonValue)

	// return failure or success
	if err != nil {
		return failure([]byte(err.Error()))
	}
	// first byte == 82
	return success(retValue)
	
}