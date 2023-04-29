package capsule

//export hostLogString
func hostLogString(messagePosition, messageLength uint32) uint32

// Log : call host function: hostLogString
// Print a string
func Log(message string) {
	messagePosition, messageSize := getStringPosSize(message)

	hostLogString(messagePosition, messageSize)
}
