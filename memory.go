package capsule

import "unsafe"

/*
func read(bufferPosition *uint32, length int) []byte {
	return readBufferFromMemory(bufferPosition, length)
}
*/


// readBufferFromMemory returns a buffer
func readBufferFromMemory(bufferPosition *uint32, length uint32) []byte {
	subjectBuffer := make([]byte, length)
	pointer := uintptr(unsafe.Pointer(bufferPosition))
	for i := 0; i < int(length); i++ {
		s := *(*int32)(unsafe.Pointer(pointer + uintptr(i)))
		subjectBuffer[i] = byte(s)
	}
	return subjectBuffer
}

// ReadBufferFromMemory returns a buffer
func ReadBufferFromMemory(bufferPosition *uint32, length uint32) []byte {
	return readBufferFromMemory(bufferPosition, length)
}

// copyBufferToMemory returns a single value (a kind of pair with position and length)
func copyBufferToMemory(buffer []byte) uint64 {
	bufferPtr := &buffer[0]
	unsafePtr := uintptr(unsafe.Pointer(bufferPtr))

	ptr := uint32(unsafePtr)
	size := uint32(len(buffer))

	return (uint64(ptr) << uint64(32)) | uint64(size)
}

// CopyBufferToMemory returns a single value
func CopyBufferToMemory(buffer []byte) uint64 {
	return copyBufferToMemory(buffer)
}


// getStringPosSize returns the memory position and size of the string
func getStringPosSize(s string) (uint32, uint32) {
	buff := []byte(s)
	ptr := &buff[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))
	return uint32(unsafePtr), uint32(len(buff))
}

// getBufferPosSize returns the memory position and size of the buffer
func getBufferPosSize(buff []byte) (uint32, uint32) {
	ptr := &buff[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))
	return uint32(unsafePtr), uint32(len(buff))
}

/* 
Allocate the in-Wasm memory region and returns its pointer to hosts.
The region is supposed to store random strings generated in hosts
*/
//export allocateBuffer
func allocateBuffer(size uint32) *byte {
	buf := make([]byte, size)
	return &buf[0]
}

