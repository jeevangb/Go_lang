package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	res := removeFile("a.txt")
	if res != nil {
		log.Println(res)
		return
	}

	fmt.Println("file removed successfully")
}

func removeFile(a string) error {
	err := os.Remove(a)
	return err
}
