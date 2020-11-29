package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	// 将第一列转为行，二分查找
	// 找到行，再在行上二分法查找
	N := len(matrix)
	if N < 1 {
		return false
	}
	temp := make([]int, N)
	for i := 0; i < N; i++ {
		if len(matrix[i]) < 1 {
			return false
		}
		temp[i] = matrix[i][0]
	}
	index, find := search(temp, 0, N-1, target)
	if find {
		return find
	}
	if index < 1 {
		return false
	}
	index--
	index, find = search(matrix[index], 0, len(matrix[index])-1, target)
	return find
}

func search(nums []int, start, end int, target int) (int, bool) {
	if start > end || start < 0 || end > len(nums) {
		return start, false
	}
	mid := (start + end) >> 1
	if nums[mid] == target {
		return start, true
	} else if nums[mid] < target {
		return search(nums, mid+1, end, target)
	}
	return search(nums, start, mid-1, target)
}

func main() {
	//fmt.Println(search([]int{1, 10, 23}, 0, 4, 6))
	matrix := make([][]int, 0)
	//matrix[0] = []int{1, 3, 5, 6}
	//matrix[1] = []int{10, 11, 16, 20}
	//matrix[2] = []int{23, 30, 34, 60}
	fmt.Println(searchMatrix(matrix, -1))
}
