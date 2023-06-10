package capsule


//export hostGetEnv
func hostGetEnv(
	messagePosition, messageLength uint32,
	returnValuePosition **uint32,
	returnValueLength *uint32) uint32

// GetEnv is an helper to use the hostGetEnv function
func GetEnv(variableName string) string {

	variableNamePosition, variableNameSize := getBufferPosSize([]byte(variableName))

	// This will be use to get the response from the host
	var responseBufferPtr *uint32
	var responseBufferSize uint32

	// Send the lessage to the host
	hostGetEnv(variableNamePosition, variableNameSize, &responseBufferPtr, &responseBufferSize)

	bufferResponseFromHost := readBufferFromMemory(responseBufferPtr, responseBufferSize)

	return string(bufferResponseFromHost)
}

/* Documentation



 */
