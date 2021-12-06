package main

import (
	"fmt"
	"runtime"
)

func main() {
	// CHANNEL
	// To transfer-receive data between goroutines
	runtime.GOMAXPROCS(2)

	var messages = make(chan string) // Make channel with `make` `chan`

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		messages <- data // This symbol <- means transfer data from the right one to the left var
	}

	// NOTE : Goroutine is asynchronous but the transfer-receive is synchronous
	// Create 3 goroutine
	go sayHelloTo("Tiara Dewangga")
	go sayHelloTo("Catherine")
	go sayHelloTo("Dominique")

	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)
	// ^ The order of the result will be different according to which goroutine has finished

	// Channel as parameter
	for _, name := range []string{"Matthew", "Mark", "John"} {
		go func(who string) {
			var data = fmt.Sprintf("helloo %s", who)
			messages <- data
		}(name)
	}

	for i := 0; i < 3; i++ {
		printChan(messages)
	}
}

func printChan(what chan string) {
	// ^ Function which has channel as parameter, using `chan`
	fmt.Println(<-what)
}
