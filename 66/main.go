package main

import "fmt"

func plusOne(digits []int) []int {
	digLen := len(digits) - 1
	digits[digLen] += 1
	for digits[digLen] > 9 {
		digits[digLen] %= 10
		digLen--
		if digLen < 0 {
			digits = append([]int{0}, digits...)
			digLen = 0
		}
		digits[digLen] += 1
	}
	return digits
}

func main() {
	fmt.Println(plusOne([]int{4,2,3,4}))
}
