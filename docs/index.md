# Capsule Module SDK

!!! info "What's new?"
    - `v0.0.4`: ✨ Add the `Success` and `Failure` functions (public functions to call `success` and `failure`) and the `StringifyHTTPResponse` function
    - `v0.0.3`: ✨ Encode `retValue.TextBody` to avoid special characters in jsonString
    - `v0.0.2`: ✨ Redis support
    - `v0.0.1`: 🎉 first release

## What is the Capsule Module SDK alias **Capsule MDK**?

Capsule MDK is a WASM SDK to develop WASM modules for the [Capsule application**s**](https://github.com/bots-garden/capsule).

> The Capsule WASM modules are developed in GoLang and compiled with **[TinyGo](https://tinygo.org/)** 💜 (with the WASI specification)

### Capsule applications?

A **Capsule** application is a **WebAssembly Module(or Function) Runner**. Right now, it exists two kind of Capsule application:

- **capsule-cli**, **CLI**. With capsule-cli, you can simply execute a **WebAssembly Capsule module** in a terminal
- **capsule-http**, an **HTTP server** that serves **WebAssembly Capsule modules**

!!! info "Good to know"
    - 🖐 you can develop your own **Capsule** application with [Capsule Host SDK](https://github.com/bots-garden/capsule-host-sdk) (alias Capsule HDK).
    - 🤗 a capsule application is **"small"** (capsule-http weighs 12M)
    - 🐳 a Capsule application is statically compiled: you can easily run it in a **Distroless** Docker container.
    - 💜 The **Capsule** applications are developed with GoLang and thanks to the **[Wazero](https://github.com/tetratelabs/wazero)** project

## What does a **WASM Capsule function** look like?

```golang
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

## What are the **added values** of a Capsule application?

A Capsule application brings superpowers to the WASM Capsule modules with **host functions**. Thanks to these **host functions**, a **WASM Capsule module** can, for example, prints a message, reads files, writes to files, makes HTTP requests, ... See the [host functions section](host-functions.md).

!!! info "Useful information for this project"
    - 🖐 Issues: [https://github.com/bots-garden/capsule-module-sdk/issues](https://github.com/bots-garden/capsule-module-sdk/issues)
    - 🚧 Milestones: [https://github.com/bots-garden/capsule-module-sdk/milestones](https://github.com/bots-garden/capsule-module-sdk/milestones)
    - 📦 Releases: [https://github.com/bots-garden/capsule-module-sdk/releases](https://github.com/bots-garden/capsule-module-sdk/releases)

