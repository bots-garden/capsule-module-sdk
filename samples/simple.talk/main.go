// Package main
package main

import (
	//"fmt"
	//"os"
	//"path/filepath"

	capsule "github.com/bots-garden/capsule-module-sdk"
)

func main() {
	capsule.SetHandle(Handle)
	
}


// Handle function
func Handle(param []byte) ([]byte, error) {


	capsule.Log("🟣 from the plugin: " + string(param))
	capsule.Print("💜 from the plugin: " + string(param))

	hostResponse := capsule.Talk([]byte("Hello I'm the WASM plugin"))
	capsule.Print("🤖" + string(hostResponse))

	return []byte("Hello " + string(param)), nil
}
