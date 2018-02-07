// Mutexes
//
// If you are working with concurrent applications, you have to deal with more
// than one resource potentially accessing some memory location. This is
// usually called race condition.
//
// In simpler terms, a race condition is similar to that moment where two
// people try to get the last piece of pizza at exactly the same time--their
// hands collide. Replace the pizza with a variable and their hands with
// Goroutines and we'll have a perfect analogy.
//
// There is one character at the dinner table to solve this issues--a father or
// mother. They have kept the pizza on a different table and we have to ask for
// permission to stand up before getting our slice of pizza. It doesn't matter
// if all the kids ask at the same time--they will only allow one kid to stand.
//
// Well, a mutex is like our parents. They'll control who can access the
// pizza--I mean, a variable--and they won't allow anyone else to access it.
//
// To use a mutex, we have to actively lock it; if it's already locked (another
// Goroutine is using it), we'll have to wait until it's unlocked again. Once we
// get access to the mutex, we can lock it again, do whatever modifications are
// needed, and unlock it again. We'll look at this using an example.

package main

import (
	"sync"
	"time"
)

// The Counter structure has a field of int type that stores the current value
// of the count. It also embeds the Mutex type from the sync package. Embedding
// this field will allow us to lock and unlock the entire structure without
// actively calling a specific field.

type Counter struct {
	sync.Mutex
	value int
}

func main() {
	counter := Counter{}

	for i := 0; i < 10; i++ {
		go func(i int) {
			counter.Lock()
			counter.value++
			defer counter.Unlock()
		}(i)
	}
	time.Sleep(time.Second)

	counter.Lock()
	defer counter.Unlock()

	println(counter.value)
}

// Checking race condition with `go run -race mutexes.go`, but comment lines 46
// and 48.
