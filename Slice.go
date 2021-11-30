package main

import "fmt"

func main() {
	// Creating Slice
	// Slice is reference of array elements

	// Initializing slice
	var fruits = []string{"apple", "mango", "banana", "watermelon"}
	// ^ Initializing slice is like array but without declaring the length
	fmt.Println("Elements of fruits: ", fruits)

	// Array is collection of values, while Slice is reference for those elements
	// Slice is reference data type
	// Slice can be made from array
	var newFruits = fruits[0:2] // Slicing the value of array by its index, starting from 0 until before index 2
	fmt.Println("Slice of newFruits: ", newFruits)

	fmt.Println("\nSlice of fruits[0:2] : ", fruits[0:2]) // [apple mango]
	fmt.Println("Slice of fruits[0:4] : ", fruits[0:4])   // [apple mango banana watermelon ]
	fmt.Println("Slice of fruits[0:0] : ", fruits[0:0])   // []
	fmt.Println("Slice of fruits[4:4] : ", fruits[4:4])   // []
	fmt.Println("Slice of fruits[4:0] :  error")          // will throw error, but I don't write the code here
	fmt.Println("Slice of fruits[:]   : ", fruits[:])     // [apple mango banana watermelon ]
	fmt.Println("Slice of fruits[2:]  : ", fruits[2:])    // [banana watermelon ]
	fmt.Println("Slice of fruits[:2]  : ", fruits[:2])    // [apple mango ]

	// Referencing from fruits
	var fruits1 = fruits[0:3]
	var fruits2 = fruits[1:4]

	var fruits11 = fruits1[1:2]
	var fruits22 = fruits2[0:1]

	fmt.Printf("\nReferencing :\n")
	fmt.Println("fruits   : ", fruits)   // [apple mango banana watermelon]
	fmt.Println("fruits1  : ", fruits1)  // [apple mango banana]
	fmt.Println("fruits2  : ", fruits2)  // [mango banana watermelon]
	fmt.Println("fruits11 : ", fruits11) // [mango]
	fmt.Println("fruits22 : ", fruits22) // [mango]

	fmt.Printf("\nChange the fruits22[0] as grape\n")
	fruits22[0] = "grape"
	fmt.Println("fruits   : ", fruits)   // [apple grape banana watermelon]
	fmt.Println("fruits1  : ", fruits1)  // [apple grape banana]
	fmt.Println("fruits2  : ", fruits2)  // [grape banana watermelon]
	fmt.Println("fruits11 : ", fruits11) // [grape]
	fmt.Println("fruits22 : ", fruits22) // [grape]
	// ^ All the mango-es are changed into grape because they have same references

	// Functions in slice
	// len() is used for knowing the length of a slice
	fmt.Println("\nLen of fruits : ", len(fruits)) // 4

	// cap() is used for knowing the maximum capacity of a slice
	fmt.Println("Cap of fruits : ", cap(fruits)) // 4

	// Let's see the difference of len() and cap()
	fmt.Println("Len of fruits[0:3] : ", len(fruits[0:3])) // 3
	fmt.Println("Cap of fruits[0:3] : ", cap(fruits[0:3])) // 4 // It has max cap 4

	fmt.Println("Len of fruits[1:4] : ", len(fruits[1:4])) // 3
	fmt.Println("Cap of fruits[1:4] : ", cap(fruits[1:4])) // 3 // Because it is started at 1, not 0

	// Append, to add value after the last index
	var appendFruits = append(fruits, "melon")
	fmt.Println("\nElements of fruits: ", fruits)            // [apple grape banana watermelon]
	fmt.Println("Elements of appendFruits: ", appendFruits)  // [apple grape banana watermelon melon]
	fmt.Println("Len of appendFruits : ", len(appendFruits)) // 5
	fmt.Println("Cap of appendFruits : ", cap(appendFruits)) // 8

	// Appending in slice is unique
	var fruits3 = fruits[0:2]
	fmt.Println("\nElements of fruits3 : ", fruits3) // [apple grape]
	fmt.Println("Len of fruits3 : ", len(fruits3))   // 2
	fmt.Println("Cap of fruits3 : ", cap(fruits3))   // 4

	var fruits4 = append(fruits3, "papaya")
	fmt.Println("\nElements of fruits : ", fruits4) // [apple grape papaya]
	fmt.Println("Elements of fruits4 : ", fruits4)  // [apple grape papaya]
	fmt.Println("Len of fruits4 : ", len(fruits4))  // 3
	fmt.Println("Cap of fruits4 : ", cap(fruits4))  // 4

	// Accessing slice with 3 indexes
	// The third index will be the capasity
	var bFruits = fruits[0:2:2]
	fmt.Println("\nSlice with 3 indexes : ")
	fmt.Println("bFruits: ", bFruits)
	fmt.Println("len(bFruits): ", len(bFruits))
	fmt.Println("cap(bFruits): ", cap(bFruits))

	// Copy Slice
	dst := make([]string, 2)
	src := []string{"a", "b", "c"}
	n := copy(dst, src)

	fmt.Println("\nCopy Slice :")
	fmt.Println(dst) // [a b] | Only contains 2 because the cap is 2
	fmt.Println(src) // [a b c]
	fmt.Println(n)   // Return the number of the copied values

	// Let's change the value
	dst = []string{"d", "e", "f"}
	src = []string{"x", "y"}
	n = copy(dst, src)
	fmt.Println("\nCopy Slice 2 :")
	fmt.Println(dst) // [x y f] | Because the len(src) < len(dst)
	fmt.Println(src) // [x y]
	fmt.Println(n)   // Return the number of the copied values

}
