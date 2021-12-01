package main

import "fmt"

func main() {
	// MAP

	// Map is like associative array or dictionary
	// Initializing map
	var fruits map[string]int
	fruits = map[string]int{} // It it a must to declare the nil to map explicitly; will throw error if this line isn't declared

	fruits["mango"] = 10
	fruits["apple"] = 12

	fmt.Println("Mango : ", fruits["mango"])
	fmt.Println("Apple : ", fruits["apple"])
	fmt.Println("Orange: ", fruits["orange"]) // 0, because wasn't initialized

	// Horizontal Style
	var days = map[string]int{"Mon": 12, "Tue": 17}
	fmt.Println(days)

	// Vertical Style
	var months = map[string]int{
		"Jan": 10,
		"Feb": 11,
		"Mar": 5, // --> Should add coma at the end of the value
	}
	fmt.Println(months)

	// Other ways to initialize map without initial value
	var days1 = map[string]int{}
	var days2 = make(map[string]int)
	var days3 = *new(map[string]int)

	fmt.Println(days1)
	fmt.Println(days2)
	fmt.Println(days3)

	// Iteration with for-range
	fmt.Printf("\nIteration with for-range : ")
	for key, val := range months {
		fmt.Println(key, " : ", val)
	}

	// Delete value in map
	delete(months, "Jan")
	fmt.Println(len(months)) // 2

	// Detection if a key in map is exist
	fmt.Println("\nDetection of a key")
	var value, isExist = months["Feb"] // "value" as the value, "isExist" as boolean if it's exist
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Item is not exist")
	}

	// Map + Slice
	fmt.Println("\nMap + Slice")
	var students = []map[string]string{
		map[string]string{"name": "Matthew", "gender": "male"},
		map[string]string{"name": "Mark", "gender": "male"},
		map[string]string{"name": "Olivia", "gender": "female"},
	}

	for _, student := range students {
		fmt.Println(student["name"], student["gender"])
	}

	// New way to use Map + Slice
	var students2 = []map[string]string{
		{"name": "Luke", "gender": "male"},
		{"name": "Catherine", "gender": "female"},
	}

	for _, student := range students2 {
		fmt.Println(student["name"], student["gender"])
	}

	// Different key in Map + Slice
	var students3 = []map[string]string{
		{"name": "John", "gender": "male"},
		{"address": "Manila", "id": "231"},
		{"class": "2B"},
	}

	for _, student := range students3 {
		fmt.Println(student)
	}
}
