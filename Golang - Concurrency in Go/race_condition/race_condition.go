// Both goroutines will access a shared variable. While the first routine waits for
// a certain time (simulating some blocking call), the second routine accesses the
// variable and writes to it. This modified the state which the first variable
// assumes after setting it. So the outcome of the print might be different than
// expected by routine one.
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go one()
	go two()

	wg.Wait()
	fmt.Println(x)
}

var x = 5

func one() {
	x = 10
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(x)

	wg.Done()
}

func two() {
	time.Sleep(500 * time.Millisecond)
	x = 6

	wg.Done()
}
