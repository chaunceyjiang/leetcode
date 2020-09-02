package main

import "fmt"

func removeElement(nums []int, val int) int {
	var tail = len(nums) - 1
	for i := 0; i <= tail; {
		if nums[i] == val {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			i++
		}
	}
	return tail+1
}

func main() {
	values := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(values, 2))
}
