// The select statement
//
// Used to handle more than one channel input within a Goroutine.
//
// In the select structure, we ask the program to choose between one or more
// channels to receive their data. We can save this data in a variable and make
// something with it before finishing the select. The select structure is just
// executed once; it doesn't matter if it is listening to more channels, it will
// be executed only once and the code will continue executing. If we want it to
// handle the same channels more than once, we have to put it in a for loop.
package main

import "time"

func sendString(ch chan<- string, s string) {
	ch <- s
}

func receiver(helloCh, goodbyeCh <-chan string, quitCh chan<- bool) {
	for {
		select {
		case msg := <-helloCh:
			println(msg)

		case msg := <-goodbyeCh:
			println(msg)

		case <-time.After(time.Second * 2):
			println("Nothing received in 2 seconds. Exiting")
			quitCh <- true
			break
		}
	}
}

func main() {
	helloCh := make(chan string, 1)
	goodbyeCh := make(chan string, 1)
	quitCh := make(chan bool)
	go receiver(helloCh, goodbyeCh, quitCh)

	go sendString(helloCh, "hello!")

	time.Sleep(time.Second)

	go sendString(goodbyeCh, "goodbye!")
	<-quitCh
}
