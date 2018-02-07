// Channels
//
// Channels are the way we communicate between processes. We could be sharing a
// memory location and using mutexes to control the processes' access. But
// channels provide us with a more natural way to handle concurrent applications
// that also produces better concurrent designs in our programs.

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	channel := make(chan string)

	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	go func() {
		channel <- "Hello World!"
		println("Finishing gorouting")
		waitGroup.Done()
	}()

	time.Sleep(time.Second)
	message := <-channel
	fmt.Println(message)
	waitGroup.Wait()
}
