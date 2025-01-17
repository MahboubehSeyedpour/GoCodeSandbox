package main

import (
	"fmt"
	"sync"
)

func main() {

	//print(numberIsEven(120)) // Output true

	//print(findTheLargestNumber(5, 3)) // Output 5

	//printArray()

	//arr := [5]int{2, 4, 8, 5, 4}
	//arrayInSlice := arr[:] // convert array to slice
	//fmt.Println(findLargestElement(arrayInSlice)) // Output: 8

	//arr := []int{2, 4, 8, 5, 4}
	//bubbleSortAsc(arr)
	//fmt.Println(arr) // Output: [2 4 4 5 8]

	//arr1 := []int{1, 3, 6, 7, 12}
	//arr2 := []int{5, 8, 9, 15}
	//fmt.Println(mergeTwoSortedArray(arr1, arr2)) // Output: [1 3 5 6 7 8 9 12 15]

	//fmt.Println(factorial(5)) // Output: 120

	//printNumbersInRange(2, 10)

	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Println(calcFactorialWithGoroutine(5, &wg))
	wg.Wait()
}
