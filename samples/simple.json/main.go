// Package main
package main

import (
	"strconv"

	capsule "github.com/bots-garden/capsule-module-sdk"
	"github.com/valyala/fastjson"
)

func main() {
	capsule.SetHandleJSON(HandleJSON)
}

// HandleJSON function
func HandleJSON(param *fastjson.Value) ([]byte, error) {

	name := param.GetStringBytes("name")
	age := param.GetInt("age")

	return []byte("Hello " + string(name) + ", age: " + strconv.Itoa(age)), nil

}
