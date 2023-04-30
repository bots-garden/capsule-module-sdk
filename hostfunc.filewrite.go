package capsule

//export hostWriteFile
func hostWriteFile(
	filePathPosition, filePathLength uint32,
	contentPosition, contentLength uint32,
	returnValuePosition **uint32, returnValueLength *uint32) uint32

// WriteFile is an helper to use the hostWriteFile function
func WriteFile(filePath string, content []byte) error {

	filePathPosition, filePathLength := getBufferPosSize([]byte(filePath))
	contentPosition, contentLength := getBufferPosSize(content)

	// This will be use to get the response from the host
	var responseBufferPtr *uint32
	var responseBufferSize uint32

	// Send the lessage to the host
	hostWriteFile(
		filePathPosition, filePathLength,
		contentPosition, contentLength,
		&responseBufferPtr, &responseBufferSize)

	bufferResponseFromHost := readBufferFromMemory(responseBufferPtr, responseBufferSize)
	
	// check if success or failure
	_, err := Result(bufferResponseFromHost)
	if err != nil {
		return err
	}
	return nil
}

/* Documentation



 */
