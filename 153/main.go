package main

import "fmt"

func findMin(nums []int) int {
	start, end := 0, len(nums)-1
	for start < end {
		mid := (end + start) / 2
		if nums[mid] < nums[start]||nums[mid]>nums[end]{
			start++
		} else {
			end = mid
		}
	}
	return nums[start]
}

func main() {
	fmt.Println(findMin([]int{3,4,5,1,2}))
}
