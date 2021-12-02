package main

import (
	"fmt"
	"strings"
)

func main() {
	// ^ Implementation of function
	var names = []string{"Katarina", "Tiara", "Dewangga"}
	printMessage("Hello", names) // Calling function

	var areaofTriangle float32
	areaofTriangle = triangleArea(10.0, 12.0) // Store the returned value to a variable
	fmt.Println("Area of triangle : ", areaofTriangle)

	var areaofTriangle2 float32
	areaofTriangle2 = triangleArea2(10.0, 12.0)
	fmt.Println("Area of triangle2 : ", areaofTriangle2)

	checkScore(91)
	checkScore(78)

	var side int = 10
	var areaOfSquare, circOfSquare = square(side) // Multiple return needs multiple variables
	fmt.Println("A square with side ", side, " has area : ", areaOfSquare, " and circumference  : ", circOfSquare)

	side = 20
	areaOfSquare, circOfSquare = square2(side)
	fmt.Println("A square with side ", side, " has area : ", areaOfSquare, " and circumference  : ", circOfSquare)
}

func printMessage(message string, arr []string) {
	// The parameters use data type
	var nameString = strings.Join(arr, " ") // String operation to join Slice with defined "glue"
	fmt.Println(message, nameString)
}

func triangleArea(base float32, height float32) float32 {
	// ^ Function with return value should add the data type of returned value
	var area = (base * height) / 2
	return area // Return the value
}

func triangleArea2(base, height float32) float32 {
	// ^ Function with some parameters which have same data type can be written in simpler way
	var area float32
	area = (base * height) / 2
	return area
}

func checkScore(score int) {
	if score >= 80 {
		fmt.Println("Good score! Great!")
		return // To return value and stop the function
		// The next line won't be processed
	}

	var neededScore int
	neededScore = 80 - score
	fmt.Println("You need ", neededScore, " to complete this course! Keep going on!")
}

func square(side int) (int, int) {
	// ^ Multiple return, write the data type in parentheses
	var area = side * side
	var circumference = 4 * side

	return area, circumference // Return with multiple variables
}

func square2(side int) (area int, circumference int) {
	// ^ Multiple return with predefined return value, the variables are declared as return value
	area = side * side
	circumference = 4 * side

	return // Just simple return to stop the function
}
