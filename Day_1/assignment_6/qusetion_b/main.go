package main

import "fmt"

func main() {

	mymap := make(map[string]int)
	mymap["apple"] = 10
	mymap["orange"] = 8
	mymap["banana"] = 5
	mymap["sapota"] = 9

	for k, v := range mymap {
		fmt.Printf("%v: %v\n", k, v)

	}
}
