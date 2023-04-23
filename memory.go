package capsule

import "unsafe"

func read(bufferPosition *uint32, length int) []byte {
	return readBufferFromMemory(bufferPosition, length)
}

// readBufferFromMemory returns a buffer from WebAssembly
func readBufferFromMemory(bufferPosition *uint32, length int) []byte {
	subjectBuffer := make([]byte, length)
	pointer := uintptr(unsafe.Pointer(bufferPosition))
	for i := 0; i < length; i++ {
		s := *(*int32)(unsafe.Pointer(pointer + uintptr(i)))
		subjectBuffer[i] = byte(s)
	}
	return subjectBuffer
}

// copyBufferToMemory returns a single value (a kind of pair with position and length)
func copyBufferToMemory(buffer []byte) uint64 {
	bufferPtr := &buffer[0]
	unsafePtr := uintptr(unsafe.Pointer(bufferPtr))

	ptr := uint32(unsafePtr)
	size := uint32(len(buffer))

	return (uint64(ptr) << uint64(32)) | uint64(size)
}
