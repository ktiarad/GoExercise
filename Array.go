package main

import "fmt"

func main() {
	// Creating Array

	// Basically, declaring the length
	var students [4]string
	students[0] = "Marc"
	students[1] = "George"
	students[2] = "Luke"
	students[3] = "Matthew"
	fmt.Println("All elements in students:", students)
	// Adding value more than the length will throw error

	// Initializing the value
	var fruits = [4]string{"apple", "mango", "jackfruit", "banana"}
	fmt.Println("All elements in fruits \t:", fruits)

	// Initializing with vertical style
	var colors = [4]string{
		"black",
		"brown",
		"blue",
		"red", // --> The last value still has to be followed by coma
	}
	fmt.Println("All elements in colors\t:", colors)

	// Initializing with "make"
	var media = make([]string, 2)
	media[0] = "Twitter"
	media[1] = "Instagram"

	// Initializing without the length
	var numbers = [...]int{2, 7, 10, 4} // Using [...] with triple dots
	fmt.Println("All elements in numbers\t:", numbers, ", length : ", len(numbers))

	// Multidimensional Array
	var numbers1 = [2][3]int{[3]int{1, 2, 3}, [3]int{4, 5, 6}}
	var numbers2 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("All elements in number1\t:", numbers1)
	fmt.Println("All elements in number2\t:", numbers2)

	// Looping through array
	fmt.Println("\nLooping through fruits:")
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Element in %d is %s\n", i, fruits[i])
	}

	// Looping with for - range
	fmt.Println("\nLooping through fruits with for-range:")
	for i, fruit := range fruits {
		// ^ i as the index, fruit as the value in every loops
		fmt.Printf("Element in %d is %s\n", i, fruit)
	}

	// Looping with _ (underscore) and for - range
	fmt.Println("\nLooping through fruits with for-range and underscore:")
	for _, fruit := range fruits {
		// ^ the index i won't be used, so just replace it with _ (underscore)
		// _ (underscore) is used when the value won't be used, it's like a black hole and will store nothing
		fmt.Printf("Current element : %s\n", fruit)
	}
}
