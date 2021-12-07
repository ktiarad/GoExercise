package main

import (
	"fmt"
	"runtime"
)

func main() {
	// BUFFERED CHANNEL
	// It is like channel, but it has the number of the buffer which determines the number of data that  can be sent at once
	// As long as the number of the sent data is smaller than the buffered channel, then it works asynchronously

	runtime.GOMAXPROCS(2)

	messages := make(chan int, 2) // The second parameter is the number of the buffer, which start from 0; 2 means 3 processes, etc.
	for i := 0; i < 5; i++ {
		// This block will be executed first because the `messages` still have no value
		fmt.Println("Send data ", i)
		messages <- i
	}
	go func() {
		for {
			i := <-messages
			fmt.Println("Receive data ", i)
		}
	}()

}
