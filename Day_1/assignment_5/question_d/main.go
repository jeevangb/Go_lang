package main

import "fmt"

func main() {
	myslice := []int{1, 2, 3, 4, 5}
	findSum(myslice)
	findAverage(myslice)

}

//sum
func findSum(myslice []int) {

	var sum int
	for i := 0; i < len(myslice); i++ {
		sum = sum + myslice[i]
	}
	fmt.Println(sum)
}

//average
func findAverage(myslice []int) {

	var sum int
	for i := 0; i < len(myslice); i++ {
		sum = sum + myslice[i]
	}
	fmt.Println(sum / len(myslice))
}
