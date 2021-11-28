package main

import "fmt"

func main() {
	// Condition statement
	// Go doesn't support ternary condition

	var score = 8

	fmt.Printf("If Else : ")
	if score == 10 {
		// ^ No need to put parantheses in condition statement
		fmt.Print("Perfect!")
	} else if score >= 8 {
		// ^ The "else if" or "else" should be put after the previous curly bracket
		fmt.Print("Great!")
	} else if score >= 5 {
		fmt.Print("Good")
	} else {
		fmt.Print("Bad")
	}

	// Temporary Variable
	var score2 = 41

	fmt.Printf("\nTemporary Variable : ")
	// The temporary variable can be used in the conditional scope only
	if avg := score2 / 5; avg == 10 {
		fmt.Print("Perfect!")
	} else if avg >= 8 {
		fmt.Print("Great!")
	} else if avg >= 5 {
		fmt.Print("Good")
	} else {
		fmt.Print("Bad")
	}

	// Switch Case
	var score3 = 6
	fmt.Printf("\nSwitch Case : ")
	switch score3 {
	case 10:
		fmt.Print("Perfect!")
	case 9:
		fmt.Print("Great!") // If a condition is fulfilled, then it won't go to the next block, as if there is "break" although it isn't
	case 5, 6, 7, 8:
		fmt.Print("Good") // Multi case, separated with coma
	default:
		{
			// The use of curly bracket is optional, but it's good better to use it especially if it contains some lines
			fmt.Print("Bad. ") // The default block will be executed if the above conditions aren't fulfilled
			fmt.Print("Keep going on!")
		}
	}

	// Switch Case with If Else Style
	var score4 = 5
	fmt.Printf("\nSwitch Case with If Else Style : ")
	switch { // The variable isn't put here
	case score4 == 10: // The if else style statement
		fmt.Print("Perfect!")
	case score4 >= 8:
		fmt.Print("Great!")
	case score4 >= 5:
		fmt.Print("Good")
	default:
		{
			fmt.Print("Bad. ")
			fmt.Print("Keep going on!")
		}
	}

	// Fallthrough
	// If the current condition is fulfilled, the fallthrough will force the next block be executed too although it doesn't fulfill the condition
	var score5 = 9
	fmt.Printf("\nFallthrough : ")
	switch { // The variable isn't put here
	case score5 == 10: // The if else style statement
		fmt.Print("Perfect!")
	case score5 >= 8:
		fmt.Print("Great!")
		fallthrough // The next block will be executed although it doesn't meet the condition
	case score5 >= 5:
		fmt.Print("Good")
	default:
		{
			fmt.Print("Bad. ")
			fmt.Print("Keep going on!")
		}
	}

}
