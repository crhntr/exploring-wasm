package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello, console!")

	alert := js.Global().Get("alert")
	alert.Invoke("Hello, global!")
}
