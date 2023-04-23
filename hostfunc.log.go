package capsule

//export hostLogString
func hostLogString(posSizePairValue uint64) uint32

// Log : call host function: hostLogString
// Print a string
func Log(message string) {
	posSizePairValue := copyBufferToMemory([]byte(message))
	hostLogString(posSizePairValue)
}
