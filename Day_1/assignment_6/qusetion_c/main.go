package main

import "fmt"

func main() {

	//map declaration
	studentDetails := make(map[string]map[string]interface{})

	//initialization here
	studentDetails["jeevan"] = map[string]interface{}{
		"Age":   20,
		"Grade": 67,
		"City":  "Bnaglore",
	}

	studentDetails["afthab"] = map[string]interface{}{
		"Age":   28,
		"Grade": 78,
		"City":  "Banglore",
	}

	studentDetails["poorvi"] = map[string]interface{}{
		"Age":   23,
		"Grade": 98,
		"City":  "mysore",
	}

	// fmt.Println(studentDetails)
	for k, v := range studentDetails {
		fmt.Println(k, v)
	}

}
