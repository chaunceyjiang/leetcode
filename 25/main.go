package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if target <= nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}
	return left
}

func main() {
	fmt.Println(searchInsert([]int{1,2,5,6},0))
}
