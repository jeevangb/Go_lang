package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("enter number")
	var num int
	fmt.Scan(&num)
	guess(num)
}

func guess(num int) {
	randNum := rand.Intn(10)
	if randNum > num {
		fmt.Println("number is too high")
	} else if randNum < num {
		fmt.Println("number is too low")
	} else {
		fmt.Println("number is equal")
	}
}
