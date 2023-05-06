// Package main
package main

import (
	capsule "github.com/bots-garden/capsule-module-sdk"
)

func main() {
	capsule.SetHandle(Handle)
}

// Handle function
func Handle(params []byte) ([]byte, error) {

	// Display the content of `params`
	capsule.Print("Module parameter(s): " + string(params))

	// Read an display an environment variable 
	capsule.Print("MESSAGE: " + capsule.GetEnv("MESSAGE"))

	// Write content to a file
	err := capsule.WriteFile("./hello.txt", []byte("👋 Hello World! 🌍"))
	if err != nil {
		capsule.Print(err.Error())
	}

	// Read content from a file
	data, err := capsule.ReadFile("./hello.txt")
	if err != nil {
		capsule.Print(err.Error())
	}
	capsule.Print("📝: " + string(data))
	

	return []byte("👋 Hello " + string(params)), nil

}

/* 
export MESSAGE="👋 Hello Capsule"
./capsule --wasm=cli-say-hello.wasm --params="Jane Doe"
*/