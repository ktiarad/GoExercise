package main

import "fmt"

func main() {
	// Looping

	// Simple For
	fmt.Println("Simple For : ")
	for i := 1; i <= 5; i++ {
		fmt.Println("Number ", i)
	}

	// For with condition only
	fmt.Println("For with condition only : ")
	var i = 1
	for i <= 5 {
		// ^ Only check the condition, like "while"
		fmt.Println("Number ", i)
		i++ // Iteration
	}

	// For without argument
	fmt.Println("For without argument : ")
	var j = 1
	for {
		fmt.Println("Number ", j)

		j++
		if j == 6 {
			break // End the looping with break in the condition statement
		}
	}

	// Continue and Break
	fmt.Println("Continue and Break : ")
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue // It will continue to the next loop
		}

		if i > 8 {
			break // It will stop the looping
		}

		fmt.Println("Number ", i)
	}

	// Label
	fmt.Println("Label :")
outerloop: // --> The label which wraps the looping
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if i == 3 {
				continue outerloop // Will go to the label "outerloop" wherever it is
			}

			fmt.Print("Matrix [", i, "][", j, "]\n")
		}
	}
}
