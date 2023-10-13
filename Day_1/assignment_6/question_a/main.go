package main

import "fmt"

func main() {
	mymap := make(map[string]float32)

	mymap["jeevan"] = 78.67
	mymap["afthab"] = 98.7
	mymap["poorvi"] = 88.67

	fmt.Println("before deletion")
	fmt.Println(mymap)

	delete(mymap, "jeevan")

	fmt.Println("After deletion")
	fmt.Println(mymap)
}
