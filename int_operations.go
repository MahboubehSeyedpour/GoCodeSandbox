package main

// Take two integers as input and print their sum.
func addToIntegers(num1 int, num2 int) int {
	return num1 + num2
}

func checkEvenOrOdd(num int) string {
	if num%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func findTheLargestOfTwoNumbers(num1 int, num2 int) int {
	if num1 >= num2 {
		return num1
	} else {
		return num2
	}
}

// check_if_a_numbers_positive_negative_or_zero
func checkIfNumberIsPositiveOrNegative(num int) string {
	if num > 0 {
		return "Positive"
	} else if num < 0 {
		return "Negative"
	} else {
		return "Zero"
	}
}

// TODO: level1: calculate_factorial
// TODO: level1: print_the_numbers_in_range(start, end)
// TODO: level2: check_if_a_numbers_prime
