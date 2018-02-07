// Buffered Channels
//
// A buffered channel works in a similar way to default unbuffered channels. You
// also pass and take values from them by using the arrows, but, unlike
// unbuffered channels, senders don't need to wait until some Goroutine picks
// the data that they are sending

package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string, 1)

	go func() {
		channel <- "Hello World! 1"
		channel <- "Hello World! 2"
		println("Finishing goroutine")
	}()

	time.Sleep(time.Second)

	message := <-channel
	fmt.Println(message)
}
