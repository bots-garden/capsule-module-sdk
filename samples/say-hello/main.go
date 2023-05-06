// Package main
package main

import (
	capsule "github.com/bots-garden/capsule-module-sdk"
)

func main() {
    // define wich function to run
	capsule.SetHandle(Handle)
}

// Handle function
func Handle(params []byte) ([]byte, error) {
    name := string(params)
    message := "👋 Hello " + name
	
	return []byte(message), nil
}
// ./capsule --wasm=say-hello.wasm --params="Bob Morane"