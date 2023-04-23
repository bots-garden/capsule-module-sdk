# Capsule Module SDK

🚧 this is a work in progress

This SDK allows to create and manage **WebAssembly modules** for host applications using the [Capsule Host Application SDK](https://github.com/bots-garden/capsule-host-sdk).

> The Capsule Host SDK use the **[Wazero](https://github.com/tetratelabs/wazero)** runtime to run the host application.

## Getting started: the capsule plugin

```golang
package main

import (
	capsule "github.com/bots-garden/capsule-module-sdk"
)

func main() {
	capsule.SetHandle(Handle)
}

// Handle function
func Handle(param []byte) ([]byte, error) {

	capsule.Log("🟣 from the plugin: " + string(param))
	capsule.Print("💜 from the plugin: " + string(param))

	return []byte("Hello " + string(param)), nil
}
```

## Getting started: the host application

👀 https://github.com/bots-garden/capsule-host-sdk#capsule-host-sdk


