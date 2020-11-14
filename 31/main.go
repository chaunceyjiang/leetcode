package main

import "fmt"

/*
１．从后向前查找第一个相邻升序的元素对 (i,j)，满足 A[i] < A[j]。此时 [j,end) 必然是降序
２．在 [j,end) 从后向前查找第一个满足 A[i] < A[k] 的 k。A[i]、A[k] 分别就是上文所说的「小数」、「大数」
３．将 A[i] 与 A[k] 交换
４．可以断定这时 [j,end) 必然是降序，逆置 [j,end)，使其升序
５．如果在步骤 1 找不到符合的相邻元素对，说明当前 [begin,end) 为一个降序顺序，则直接跳到步骤 4
*/

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	//	find 第一个相邻升序 A[i] <A[j]
	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	if i >= 0 { // 不是最大序列
		// 查找第一个满足 A[i] < A[k]
		for nums[i] >= nums[k] {
			k--
		}
		// 将 A[i] 与 A[k] 交换
		nums[i], nums[k] = nums[k], nums[i]
	}
	for i, j := j, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func main() {
	a := []int{4, 5, 3, 2, 1}
	nextPermutation(a)
	fmt.Println(a)
}
