package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	// GOROUTINE
	// Mini by thread which works asynchronously, one of the most important thing in concurrent programming in Go
	runtime.GOMAXPROCS(2)

	go print(100, "hello") // Create new goroutine with `go` keyword
	print(100, "what's up?")
	// ^ It will print the "hello" and "what's" up by turns, not in order

	var input string
	fmt.Scanln(&input) // By pressing any key, it will end the program
	// ^ It will do blocking, because probably the time needed to execute goroutine print() will be longer than the main goroutine `main()`
}
