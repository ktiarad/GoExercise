package main

import (
	"fmt"
	"strings"
)

func main() {
	// VARIADIC FUNCTION
	// A function which has "unlimited" parameters
	var avg float64 = getAvg(3, 5, 10, 2, 1, 9, 12, 11, 7, 21, 51)
	fmt.Printf("Average : %.2f\n", avg)

	// Variadic with slice
	var numbers = []int{1, 21, 51, 52, 53, 11}
	avg = getAvg(numbers...) // The parameter is the variable followed by ...
	fmt.Printf("Average : %.2f\n", avg)

	// Variadic with variable and slice as paramters
	var name string = "Tiara"
	var hobbies = []string{"coding", "digital art", "reading book"}
	myHobbies(name, hobbies...)                                  // The variadic param should be the last param
	myHobbies("Dewangga", "playing game", "surfing on internet") // Another way to put variadic param

	// CLOSURE FUNCTION
	// An anonymous function; can be stored in a variable
	var getSumAvg = func(numbers []int) (int, float32) {
		// ^ A function with name and stored in variable
		var sum int = 0
		var avg float32

		for _, num := range numbers {
			sum += num
		}

		avg = float32(sum) / float32(len(numbers))

		return sum, avg
	}
	var sum2, avg2 = getSumAvg(numbers)
	fmt.Printf("The total number is %d and the average is %.2f\n", sum2, avg2)

	// Immediately-Invoked Function Expression (IIFE)
	// ^ The value is declared first and immediately processed; for one-time used only
	var scores = []int{90, 78, 56, 88, 80, 79, 75, 81}
	var goodScore = func(minScore int) []int {
		var goodScores = []int{}

		for _, score := range scores {
			if score < minScore {
				continue
			}
			goodScores = append(goodScores, score)
		}
		return goodScores
	}(75) // The parameter is put in the parentheses at the end of the program

	fmt.Println("The good scores are : ", goodScore)

	// Closure as return
	var sum, sum3 = simpleSum(numbers)
	fmt.Printf("The values are %d and %d\n", sum, sum3)

	// FUNCTION AS PARAMETER
	var names = []string{"james", "jason", "ronald"}
	var dataContainsO = filter(names, func(check string) bool {
		return strings.Contains(check, "o") // strings.Contains check if a value has certain value
	})
	fmt.Println("Real data : ", names)
	fmt.Println("Data contains O : ", dataContainsO)
}

func getAvg(numbers ...int) float64 {
	var total int = 0

	for _, number := range numbers {
		total += number // Equals to : total = total + number
	}

	var avg float64 = float64(total) / float64(len(numbers)) // Use float64() to convert it into float
	return avg

}

func myHobbies(name string, hobbies ...string) {
	var hobbiesStr = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, I'm %s, my hobbies are %s.\n", name, hobbiesStr)
}

func simpleSum(numbers []int) (int, int) {
	var sum int = 0

	for _, n := range numbers {
		sum += n
	}

	return sum, 2
	// The second returned value should be
	// func() int { return 2 } // as closure function for returned value, but it doesn't work ._.
}

func filter(data []string, callback func(string) bool) []string {
	// ^ "callback" is the closure function with data type func(string) bool
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

// Alias
type FilterCallback func(string) bool // Using type
func filter2(data []string, callback FilterCallback) []string {
	// The scheme `func(string) bool` is changed with FilterCallack
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
