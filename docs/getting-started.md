# 🚀 Getting Started

## Create a Capsule module

Create a directory `say-hello`

```bash
mkdir say-hello
cd say-hello
```

Initialize a new project in `say-hello`:

```bash
go mod init say-hello
```

Install the Capsule MDK dependencies:
```bash
go get github.com/bots-garden/capsule-module-sdk
```


Create a new file `main.go` in `say-hello`:

```go
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
```

## Build the WASM Capsule module

```bash
tinygo build -o say-hello.wasm -scheduler=none --no-debug -target wasi ./main.go
```

## Execute the WASM Capsule module

You need to download the last capsule CLI: [https://github.com/bots-garden/capsule/releases](https://github.com/bots-garden/capsule/releases)

```bash
./capsule --wasm=say-hello.wasm --params="Bob Morane"
```

You should get: `👋 Hello Bob Morane`
