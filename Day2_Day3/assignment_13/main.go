package main

import (
	"fmt"
	"os"
	"strings"
)

func createFile() (*os.File, error) {
	//create file ,if not possible throw to caller
	f, err := os.Create("hello.txt")
	if err != nil {
		return nil, err
	}

	data := []byte("   ")
	f.Write(data)

	return f, nil
}

func openFile() (int, error) {
	// f, err := os.Open("hello.txt")
	// if err != nil {
	// 	return 0, err
	// }

	//container
	// data := make([]byte, 100)

	// //try read data and store in byte format
	// _, err = f.Rea(data)

	data, err := os.ReadFile("hello.txt")
	if err != nil {

		return 0, err

	}

	//print data present inside the file
	//convert the byte format of data to string type
	words := string(data)
	fmt.Println(words)

	// wordtrim := strings.Trim(words, " ")

	//it store data in slice format and  its splits at every whitw space
	content := strings.Fields(words)

	res := len(content)
	return res, nil

}
func removefile(s string) error {
	return os.Remove(s)

}

func main() {
	//call to create ,else handle it
	_, err := createFile()
	if err != nil {
		fmt.Println("not able to create")
		return
	}
	//try to open and read ,else handle it
	out, err := openFile()
	if err != nil {
		fmt.Println("There is some error in data format")
		return
	}
	// fmt.Println(out)
	//try to count number of words in a file
	// if out == 0 {
	// 	fmt.Println("file is empty")
	// } else {
	fmt.Println("number of words is : ", out)
	// }

	// error := removefile("hello.txt")
	// if error != nil {
	// 	fmt.Println("File removed succesfully")
	// 	return
	// }
	// fmt.Println("file not preset")

}
