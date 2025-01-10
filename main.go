package main

import "fmt"

func main() {

	//print(addToIntegers(2, 3)) // Output 5

	//print(checkEvenOrOdd(120)) // Output Even

	//print(findTheLargestOfTwoNumbers(5, 3)) // Output 5

	//printArray()

	//arr := [5]int{2, 4, 8, 5, 4}
	//arrayInSlice := arr[:] // convert array to slice
	//fmt.Println(findLargestElement(arrayInSlice)) // Output: 8

	//arr := []int{2, 4, 8, 5, 4}
	//bubbleSortAsc(arr)
	//fmt.Println(arr)

	arr1 := []int{1, 3, 6, 7, 12}
	arr2 := []int{5, 8, 9, 15}
	fmt.Println(mergeTwoSortedArray(arr1, arr2))
}
