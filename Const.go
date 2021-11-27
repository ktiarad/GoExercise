package main

import "fmt"

func main() {
	// Declaring consts
	const firstName string = "Tiara"

	const lastName = "Dewangga" // --> Type Inference in const

	fmt.Printf("Hello, %s %s!\n", firstName, lastName)

	// Some variants of Print
	fmt.Println("Tiara Dewangga")    // Automatically adding enter
	fmt.Println("Tiara", "Dewangga") // Automatically adding enter and space between string in params

	fmt.Print("Tiara Dewangga\n")     // Should add \n
	fmt.Print("Tiara", " Dewangga\n") // Should add space because it doesn't add space automatically
	fmt.Print("Tiara", " ", "Dewangga\n")
}
