package main

import "fmt"

func main() {
	// Pointer is reference or address of memory
	// Pointer variable is the memory address of a value

	var numberA int = 4
	// Initalizing of pointer variable
	var numberB *int = &numberA
	// ^ Using asterix (*) to initialize pointer variable
	// ^ Using ampersand (&) to get the value of a pointer

	fmt.Println("Number A (value)   : ", numberA)  // 4
	fmt.Println("Number A (address) : ", &numberA) // 0xc000012088 ; using ampersand (&) to get the pointer of the ordinary variable; called referencing
	fmt.Println("Number B (value)   : ", *numberB) // 4 ; using asterix (*) to get the value of the pointer variable
	fmt.Println("Number B (address) : ", numberB)  // 0xc000012088
	fmt.Println("")

	// Changing the value of the pointer will affect the other variables which reference to this memory
	numberA = 21
	fmt.Println("Number A (value)   : ", numberA)  // 21
	fmt.Println("Number A (address) : ", &numberA) // 0xc000012088
	fmt.Println("Number B (value)   : ", *numberB) // 21
	fmt.Println("Number B (address) : ", numberB)  // 0xc000012088

	// Pointer as parameter
	fmt.Println("")
	fmt.Println("Before : ", numberA) //21
	change(&numberA, 51)
	fmt.Println("After  : ", numberA)  //51
	fmt.Println("After  : ", *numberB) //51

}

func change(original *int, new int) {
	*original = new
}
