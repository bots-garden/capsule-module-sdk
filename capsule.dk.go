// Package capsule SDK for WASM plugin
package capsule

import (
	"errors"
)

const isFailure = rune('F')
const isSuccess = rune('S')

func success(buffer []byte) uint64 {
	return copyBufferToMemory(append([]byte(string(isSuccess)), buffer...))
}

func failure(buffer []byte) uint64 {
	return copyBufferToMemory(append([]byte(string(isFailure)), buffer...))
}

// Success function
func Success(buffer []byte) uint64 {
	return success(buffer)
}

// Failure function
func Failure(buffer []byte) uint64 {
	return failure(buffer)
}

// Result function
func Result(data []byte,) ([]byte, error) {
	if data[0] == byte(isSuccess) {
		return data[1:], nil
	}
	return nil, errors.New(string(data[1:]))
}
