package main

import "fmt"

func main() {
	
	// strings
	// var nameOne string = "mario" // explicit type declaration
	// var nameTwo = "luigi" // type inference
	// var nameThree string // no value assigned

	// fmt.Println(nameOne, nameTwo, nameThree)
	
	// nameOne = "peach"
	// nameThree = "bowser"
	
	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameFour := "yoshi" // shorthand declaration, colon is only used the first time for initialization and only in functions
	// fmt.Println(nameFour)

	// integers
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40 // all valid

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory

	// var numOne int8 = 1225  // not valid, as int8 is only from -128 to 127
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint8 = 255

	// float

	var scoreOne float32 = 25.98
	var scoreTwo float64 = 88312309823092183092893.7 
	scoreThree := 1.5
}