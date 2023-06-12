# 🧰 Helpers

> 🚧 this is a work in progress

## Success and Failure
> introduced in v0.0.4

`Success` and `Failure` copy an `[]byte` to the shared memory, and return the position and the length of the value "packed" into a single value (`uint64`).

If you use `Success`, the `[]byte` is prefixed by `rune('S')` before being copied to the shared memory.

If you use `Failure`, the `[]byte` is prefixed by `rune('F')` before being copied to the shared memory.

This value will be used by the function `CallHandleFunction` of the [Host SDK](https://github.com/bots-garden/capsule-host-sdk/blob/main/runtime.go) that will use the [`Result` function](https://github.com/bots-garden/capsule-host-sdk/blob/main/capsule.dk.go) to extract the result status (success or failure) and the result value (value or error).

```golang
// Package main
package main

import (
	"strconv"
	"github.com/bots-garden/capsule-module-sdk"
)

// OnHealthCheck function
//export OnHealthCheck
func OnHealthCheck() uint64 {
	capsule.Print("⛑️ OnHealthCheck")

	response := capsule.HTTPResponse{
		JSONBody: `{"message": "OK"}`,
		Headers: `{"Content-Type": "application/json; charset=utf-8"}`,
		StatusCode: 200,
	}

	return capsule.Success([]byte(capsule.StringifyHTTPResponse(response)))
}

func main() {
	capsule.SetHandleHTTP(func (param capsule.HTTPRequest) (capsule.HTTPResponse, error) {
		return capsule.HTTPResponse{
			TextBody: "👋 Hey",
			Headers: `{"Content-Type": "text/plain; charset=utf-8"}`,
			StatusCode: 200,
		}, nil
		
	})
}
```
> 👋 don't forget to export the `OnHealthCheck` function

## StringifyHTTPResponse
> introduced in v0.0.4

`StringifyHTTPResponse` converts a `capsule.HTTPResponse` into a string.

```golang
response := capsule.HTTPResponse{
	JSONBody: `{"message": "OK"}`,
	Headers: `{"Content-Type": "application/json; charset=utf-8"}`,
	StatusCode: 200,
}

str := capsule.StringifyHTTPResponse(response)
```