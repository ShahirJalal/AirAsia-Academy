package main

import "fmt"

func main() {
	welcome()
	addition(5, 3)
	result := multiply(4, 10)
	fmt.Println(result)
	sum, multiple := maths(4, 4)
	fmt.Println(sum, " and ", multiple)

	sumNew, multipleNew := mathsNew(8, 2)
	fmt.Println(sumNew, " and ", multipleNew)
}

func mathsNew(i1, i2 int) {
	s = i1 + 
}

func maths(i1 int, i2 int) (int, int) {
	return i1 + i2, i1 * i2
}

func multiply(i1, i2 int) int {
	return i1 * i2
}

func addition(i1, i2 int) {
	fmt.Println(i1 + i2)
}

func welcome() {
	fmt.Println("Welcome tp functions in Go Lang")
}
