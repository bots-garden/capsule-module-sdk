package capsule

//export hostReadFile
func hostReadFile(
	filePathPosition, filePathLength uint32,
	returnValuePosition **uint32, returnValueLength *uint32) uint32

// ReadFile is an helper to use the hostReadFile function
func ReadFile(filePath string) ([]byte, error) {

	filePathPosition, filePathLength := getBufferPosSize([]byte(filePath))

	// This will be use to get the response from the host
	var responseBufferPtr *uint32
	var responseBufferSize uint32

	// Send the message to the host
	hostReadFile(
		filePathPosition, filePathLength,
		&responseBufferPtr, &responseBufferSize)

	bufferResponseFromHost := readBufferFromMemory(responseBufferPtr, responseBufferSize)

	// check if success or failure
	data, err := Result(bufferResponseFromHost)
	if err != nil {
		return nil, err
	}
	return data, nil
}

/* Documentation

 */
