package capsule

var handleFunction func(param []byte) ([]byte, error)

// SetHandle sets the handle function
func SetHandle(function func(param []byte) ([]byte, error)) {
	handleFunction = function
}

//export callHandle
func callHandle(subjectPosition *uint32, length uint32) uint64 {
	// read the memory to get the parameter
	subjectBytes := readBufferFromMemory(subjectPosition, length)

	// call the handle function
	retValue, err := handleFunction(subjectBytes)

	// return failure or success
	if err != nil {
		return failure([]byte(err.Error()))
	}
	// first byte == 82
	return success(retValue)
	
}