package main

import "fmt"

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := (start + end) / 2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			start = mid + 1
		case nums[mid] > target:
			end = mid - 1
		}

	}
	return -1
}

func main() {
		fmt.Println(search([]int{5},5))
}
