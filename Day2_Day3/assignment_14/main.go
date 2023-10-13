package main

import (
	"fmt"
	"runtime/debug"
)

func divide(a, b int) {
	defer errorRecobvery()

	res := a / b

	fmt.Println("result: ", res)

}
func main() {
	divide(12, 0)
}
func errorRecobvery() {
	msg := recover()
	if msg != nil {
		fmt.Println(string(debug.Stack()))
		fmt.Println(msg)
		fmt.Println("panic is recovered...")
	}

}
