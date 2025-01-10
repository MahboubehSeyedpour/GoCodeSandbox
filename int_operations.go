package main

import (
	"fmt"
)

func numberIsEven(value int) bool {
	return value%2 == 0
}

func findTheLargestNumber(num1 int, num2 int) int {
	return max(num1, num2)
}

// calculate_factorial
func factorial(num int) int {
	if num == 1 || num == 0 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}

// print_the_numbers_in_range(start, end)
// pre-assumption that end is greater that start
func printNumbersInRange(start int, end int) {
	for i := start; i < end; i++ {
		fmt.Println(i)
	}
}

// TODO: level2: check_if_a_numbers_prime
