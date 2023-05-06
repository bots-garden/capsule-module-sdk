# Developer Guide

## First HTTP Module


Create a directory `http-say-hello`

```bash
mkdir http-say-hello
cd http-say-hello
```

Initialize a new project in `http-say-hello`:

```bash
go mod init http-say-hello
```

Install the Capsule MDK dependencies:
```bash
go get github.com/bots-garden/capsule-module-sdk
```

Create a new file `main.go` in `http-say-hello`:

```go
// Package main
package main

import (
	"strconv"

	"github.com/bots-garden/capsule-module-sdk"
	"github.com/valyala/fastjson"
)

func main() {
	capsule.SetHandleHTTP(Handle)
}

// Handle function 
func Handle(param capsule.HTTPRequest) (capsule.HTTPResponse, error) {
	
	capsule.Print("📝: " + param.Body)
	capsule.Print("🔠: " + param.Method)
	capsule.Print("🌍: " + param.URI)
	capsule.Print("👒: " + param.Headers)
	
	var p fastjson.Parser
	jsonBody, err := p.Parse(param.Body)
	if err != nil {
		capsule.Log(err.Error())
	}
	message := string(jsonBody.GetStringBytes("name")) + " " + strconv.Itoa(jsonBody.GetInt("age"))
	capsule.Log(message)

	response := capsule.HTTPResponse{
		JSONBody: `{"message": "`+message+`"}`,
		Headers: `{"Content-Type": "application/json; charset=utf-8"}`,
		StatusCode: 200,
	}

	return response, nil
}
```
> - `capsule.SetHandleHTTP(Handle)` defines the called wasm function
> - `capsule.Print()` and `capsule.Log()` are host functions defined in the Capsule **HDK** [https://github.com/bots-garden/capsule-host-sdk](https://github.com/bots-garden/capsule-host-sdk)
> - `capsule.HTTPRequest` and `capsule.HTTPResponse` are structures defined in [models.go](https://github.com/bots-garden/capsule-module-sdk/blob/main/models.go)

Build the wasm module:
```bash
tinygo build -o http-say-hello.wasm -scheduler=none --no-debug -target wasi ./main.go
```

**Serve** the module:
> You need to download the last **capsule-http** runner: [https://github.com/bots-garden/capsule/releases](https://github.com/bots-garden/capsule/releases)
```bash
./capsule-http --wasm=http-say-hello.wasm --httpPort=8080
```
> - `--wasm` flag: the path to the wasm file
> - `--httpPort` flag: the HTTP port to listen on

**Call** the module (function):

```bash
curl -X POST http://localhost:8080 \
    -H 'Content-Type: application/json; charset=utf-8' \
    -d '{"name":"Bob Morane","age":42}'
```

*output:* (curl response)
```bash
{"message":"Bob Morane 42"}
```

*output:* (on the capsule-http side)
```bash
📝: {"name":"Bob Morane","age":42}
🔠: POST
🌍: http://localhost:8080/
👒: "Content-Type":"application/json; charset=utf-8","User-Agent":"curl/7.81.0","Accept":"*/*","Host":"localhost:8080","Content-Length":"30"
2023-05-06 08:55:55.717252231 +0200 CEST m=+24.164260454 : Bob Morane 42
```

> You can find more samples on [/capsule/capsule-http](https://github.com/bots-garden/capsule/capsule-http)

