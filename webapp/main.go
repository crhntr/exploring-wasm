package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello, world! I am wasm!")

	alert := js.Global().Get("alert")
	alert.Invoke("Hello, globe!")
}
