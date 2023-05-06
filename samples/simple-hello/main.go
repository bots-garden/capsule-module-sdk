// 🚧 THIS IS A WORK IN PROGRESS 🚧

// Package main
package main

import (
	capsule "github.com/bots-garden/capsule-module-sdk"
)

func main() {
	capsule.SetHandle(Handle)
}

// Host function usage

//export hostPrintHello
func hostPrintHello(posSizePairValue uint64) uint32

// PrintHello a string
func PrintHello(message string) {
	posSizePairValue := capsule.CopyBufferToMemory([]byte(message))
	hostPrintHello(posSizePairValue)
}

// Handle function
func Handle(param []byte) ([]byte, error) {

	PrintHello(string(param))

	capsule.Log("🟣 from the plugin: " + string(param))
	capsule.Print("💜 from the plugin: " + string(param))

	return []byte("Hello " + string(param)), nil
}
