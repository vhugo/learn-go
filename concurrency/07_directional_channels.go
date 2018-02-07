// Directional channels
//
// One cool feature about Go channels is that, when we use them as parameters,
// we can restrict their directionality so that they can be used only to send or
// to receive. The compiler will complain if a channel is used in the restricted
// direction. This feature applies a new level of static typing to Go apps and
// makes code more understandable and more readable.

package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string, 1)

	// The line where we launch the new Goroutine go func(ch chan<- string) states
	// that the channel passed to this function can only be used as an input
	// channel, and you can't listen to it.
	go func(ch chan<- string) {
		ch <- "Hello World!"
		println("Finishing goroutine")
	}(channel)

	time.Sleep(time.Second)

	message := <-channel
	fmt.Println(message)
}

// We can also pass a channel that will be used as a receiver channel only
func receivingChannel(ch <-chan string) {
	msg := <-ch
	println(msg)
	// ch <- "hello"
	// invalid operation: ch <- "hello" (send to receive-only type <-chan string)
}

// As you can see, the arrow is on the opposite side of the keyword chan,
// indicating an extracting operation from the channel. Keep in mind that the
// channel arrow always points left, to indicate a receiving channel, it must go
// on the left, and to indicate an inserting channel, it must go on the right.
