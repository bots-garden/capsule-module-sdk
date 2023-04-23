// Package main
package main

import (
	"fmt"
	"os"
	"path/filepath"

	capsule "github.com/bots-garden/capsule-module-sdk"
)

func main() {
	if filepath.Ext(os.Args[0]) == ".wasm" {
		// CLI mode
		value, err := Handle([]byte(os.Args[1]))
		fmt.Println(string(value), err)
	} else {
		// Plugin mode
		capsule.SetHandle(Handle)
		
	}
}

// Handle function
func Handle(param []byte) ([]byte, error) {

	return []byte("Hello " + string(param)), nil

}
