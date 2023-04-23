package capsule

//export hostPrintString
func hostPrintString(posSizePairValue uint64) uint32

// Print : call host function: hostPrintString
// Print a string
func Print(message string) {
	posSizePairValue := copyBufferToMemory([]byte(message))
	hostPrintString(posSizePairValue)
}
