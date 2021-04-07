package main

import "fmt"

func search(nums []int, target int) bool {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := (end + start) / 2
		if nums[mid] == target {
			return true
		}
		if nums[start] == nums[mid] {
			//　连续相等
			start++
			continue
		}
		if nums[start] < nums[mid] {
			if nums[mid] > target && nums[start] <= target {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if nums[mid] < target && nums[end] >= target {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return false
}
func main() {
	fmt.Println(search([]int{1,1,1,1,1,1,1,1,1,13,1,1,1,1,1,1,1,1,1,1,1,1}, 13))
}
