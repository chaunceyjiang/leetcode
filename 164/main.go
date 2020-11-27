package main

import "fmt"

func maximumGap(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	quicksort(nums, 0, len(nums)-1)
	max := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] > max {
			max = nums[i+1] - nums[i]
		}
	}
	return max
}

func quicksort(nums []int, start, end int) {
	if start >= end {
		return
	}
	t := nums[start]
	i, j := start, end
	for i < j {
		for i < j && nums[j] > t {
			// 在当前值后面的数，比当前值大，则不需要任何处理
			j--
		}
		for i < j && nums[i] <= t {
			// 在当前值前面的数，比当前值小，则不需要任何处理
			i++
		}
		// 找到一个位置错误的数据，则进行交换
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[start], nums[i] = nums[i], nums[start]
	quicksort(nums, start, i-1)
	quicksort(nums, i+1, end)
}

func main() {
	a := []int{100,3,2,1}
	fmt.Println(maximumGap(a))
}
