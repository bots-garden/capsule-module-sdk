// Package capsule SDK for WASM plugin
package capsule

const isFailure = rune('F')
const isSuccess = rune('S')


func success(buffer []byte) uint64 {
	return copyBufferToMemory(append([]byte(string(isSuccess)), buffer...))
}

func failure(buffer []byte) uint64 {
	return copyBufferToMemory(append([]byte(string(isFailure)), buffer...))
}

