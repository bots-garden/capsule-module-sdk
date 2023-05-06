# Developer Guide

## First CLI Module

Create a directory `cli-say-hello`

```bash
mkdir cli-say-hello
cd cli-say-hello
```

Initialize a new project in `cli-say-hello`:

```bash
go mod init cli-say-hello
```

Install the Capsule MDK dependencies:
```bash
go https://github.com/bots-garden/capsule-module-sdk
```

Create a new file `main.go` in `cli-say-hello`:

```go
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
```
> - `capsule.SetHandle(Handle)` defines the called wasm function
> - `capsule.Print()`, `capsule.GetEnv()`, `capsule.WriteFile()` and `capsule.ReadFile()` are host functions defined in the Capsule **HDK** [https://github.com/bots-garden/capsule-host-sdk](https://github.com/bots-garden/capsule-host-sdk)

Build the wasm module:
```bash
tinygo build -o cli-say-hello.wasm -scheduler=none --no-debug -target wasi ./main.go
```

Run the module:
> You need to download the last capsule CLI: [https://github.com/bots-garden/capsule/releases](https://github.com/bots-garden/capsule/releases)
```bash
export MESSAGE="👋 Hello Capsule"
./capsule --wasm=cli-say-hello.wasm --params="Jane Doe"
```
> - `--wasm` flag: the path to the wasm file
> - `--mode` flag: the parameter to pass to the wasm module


*output:*
```bash
Module parameter(s): Jane Doe
MESSAGE: 👋 Hello Capsule
📝: 👋 Hello World! 🌍
👋 Hello Jane Doe
```

> You can find more samples on [/capsule/capsule-cli](https://github.com/bots-garden/capsule/capsule-cli)
