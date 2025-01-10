package main

import "fmt"

// Create an array of integers, populate it with 5 numbers, and print each number.
func printArray() {
	arr := [5]int{1, 2, 3, 4, 5}
	//fmt.Println(arr) // simple print, Output [1,2,3,4,5]
	for index, item := range arr {
		fmt.Printf("index: %d, value: %d\n", index, item)
	}
}

func findLargestElement(arr []int) int {
	max := arr[0]
	for _, item := range arr {
		if max < item {
			max = item
		}
	}
	return max
}

// Sort an Array (Implement a sorting algorithm (e.g., bubble sort) to sort an array of integers in ascending order.)
func bubbleSortAsc(arr []int) {

	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

}

// Merge Two Sorted Arrays (Write a program to merge two sorted arrays into one sorted array without using additional libraries.)
// The assumption is that both arrays are sorted in the same order and are both ascending or both descending.
func mergeTwoSortedArray(arr1 []int, arr2 []int) []int {

	result := make([]int, 0, len(arr1)+len(arr2)) // create new slice with length = 0 and capacity = len(arr1)+len(arr2)

	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	// Append remaining elements from arr1, if any
	for i < len(arr1) {
		result = append(result, arr1[i])
		i++
	}

	// Append remaining elements from arr2, if any
	for j < len(arr2) {
		result = append(result, arr2[j])
		j++
	}

	return result
}

// TODO level 3: Write a program to merge two sorted arrays into one sorted array without using additional libraries.
// TODO level 3: Given an array and a target sum, find all subarrays whose elements sum up to the target.
// TODO level 4: Write a program to find and print all unique elements in an array.
