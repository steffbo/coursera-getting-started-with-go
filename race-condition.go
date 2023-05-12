/*
Write two goroutines which have a race condition when executed concurrently.

Explain what the race condition is and how it can occur.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	x := 0

	// x is shared in both goroutines and the order of the operations is non-deterministic
	for i := 0; i < 10; i++ {
		// x could be incremented more than once before it gets doubled leading to different values
		go inc(&x)
		go double(&x)
	}

	// wait for the loop to finish
	time.Sleep(time.Duration(100))

	// result will very likely differ between runs
	fmt.Println(x)
}

func inc(x *int) {
	*x++
}

func double(x *int) {
	*x = *x * 2
}
