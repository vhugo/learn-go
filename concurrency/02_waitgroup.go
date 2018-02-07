// WaitGroup synchronize many concurrent Goroutines. Whenever we have to wait
// for one Goroutine to finish, we add 1 to the group, and once all of them
// are	added, we ask the group to wait. When the Goroutine finishes, it says
// Done and the WaitGroup will take one from the group.

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup

	goRoutines := 5
	wait.Add(goRoutines)

	for i := 0; i < goRoutines; i++ {
		go func(goRoutineID int) {
			fmt.Printf("ID:%d: Hello goroutines!\n", goRoutineID)
			wait.Done() // wait.Add(-1)
		}(i)
	}

	wait.Wait()
}
