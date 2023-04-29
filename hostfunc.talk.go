package capsule


//export hostTalk
func hostTalk(messagePosition, messageLength uint32, returnValuePosition **uint32, returnValueLength *uint32) uint32

// Talk is an helper to use the hostTalk function
func Talk(bufferMessageToHost []byte) []byte {

	messagePosition, messageSize := getBufferPosSize(bufferMessageToHost)
	
	// This will be use to get the response from the host
	var responseBufferPtr *uint32
	var responseBufferSize uint32

	// Send the lessage to the host
	hostTalk(messagePosition, messageSize, &responseBufferPtr, &responseBufferSize)

	bufferResponseFromHost := readBufferFromMemory(responseBufferPtr, responseBufferSize)

	return bufferResponseFromHost
}
