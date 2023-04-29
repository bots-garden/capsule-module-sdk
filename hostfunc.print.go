package capsule

//export hostPrintString
func hostPrintString(messagePosition, messageLength uint32) uint32

// Print : call host function: hostPrintString
// Print a string
func Print(message string) {
	messagePosition, messageSize := getStringPosSize(message)

	hostPrintString(messagePosition, messageSize)
}
