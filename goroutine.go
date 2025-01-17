package main

import "sync"

// TODO Calculate Factorial with Goroutines
func calcFactorialWithGoroutine(start int, wg *sync.WaitGroup) int {

	defer wg.Done()

	if start == 0 || start == 1 {
		return 1
	} else {
		wg.Add(1)
		return start * calcFactorialWithGoroutine(start-1, wg)
	}

}
