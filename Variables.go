package main

import "fmt"

func main() {
	// Manifest Typing : declare variable with data type
	var firstName string = "Tiara" // --> Using data type with value initalized

	var lastName string
	lastName = "Dewangga" // --> Declare the variable first then initialize the value

	fmt.Printf("Halo, %s %s\n", firstName, lastName) // --> The printing is like C, %s for string

	// Type Inference : declare variable without data type
	hobi := "Ngoding" // --> Without "var" and data type
	fmt.Printf("Hobinya : %s\n", hobi)

	// Declaration multi variable
	var a, b, c string
	a, b, c = "a", "b", "c"
	fmt.Printf("Value of a is %s, b is %s, c is %s\n", a, b, c)

	var d, e, f string = "d", "e", "f"
	fmt.Printf("Value of d is %s, e is %s, f is %s\n", d, e, f)

	g, h, i := "g", "h", "i" // --> With Type Inference
	fmt.Printf("Value of g is %s, h is %s, i is %s\n", g, h, i)

	name, age, isSunday, liter := "Sasha", 12, true, 9.2 // --> Multi data type
	fmt.Printf("Value of name is %s, age is %d, isSunday is %t, liter is %.2f\n", name, age, isSunday, liter)
}
