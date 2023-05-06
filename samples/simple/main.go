// Package main
package main

import (
	capsule "github.com/bots-garden/capsule-module-sdk"
)

func main() {
	capsule.SetHandle(Handle)
}

// Handle function
func Handle(param []byte) ([]byte, error) {

	return []byte("Hello " + string(param)), nil
}
