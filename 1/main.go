package main

import "fmt"

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, ok := m[target-nums[i]]
		if !ok {
			m[nums[i]] = i
		} else {
			res[0] = m[target-nums[i]]
			res[1] = i
			break
		}
	}
	return res
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15},18))
}
