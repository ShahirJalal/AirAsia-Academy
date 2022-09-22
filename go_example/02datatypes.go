package main

import "fmt"

func main() {
	fmt.Println("Datatypes")
	// declare a variable in go
	// var variableName datatype
	var score int
	var cityName string
	var check bool
	var lifeLine float32
	var autoScore string


	score = 88
	cityName = "KL"
	check = true
	lifeLine = 88.88
	

	fmt.Println(score)
	fmt.Println(cityName)
	fmt.Println(check)

	fmt.Println("Score:", score)
	fmt.Println("City name:", cityName)
	fmt.Println("Check:", check)
	fmt.Println("lifeLine:", lifeLine)
	fmt.Println("My score is ", score, ". I live in ", cityName)

	//use format specifier
	//add backspace \n to add a new line
	fmt.Printf("My score is %v. I live in %v. \n", score, cityName)
	fmt.Printf("My lifeline is %v.", lifeLine)

	//%T format specifier -> get datatype of variable.
	fmt.Printf("Datatype of lifeline variable is %v \n", lifeLine)

	//auto type inference
	//let Go detect the datatype of the variable automatically
	// use :=
	autoScore = "PASS"
	fmt.Printf("My autoScore is %v. \n", autoScore)
	fmt.Printf("Datatype of autoScore variable is %v \n", autoScore)

}