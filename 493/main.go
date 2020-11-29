package main

import "fmt"

func reversePairs(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	// 递归 拷贝新数组
	n1 := append([]int{}, nums[:n/2]...)
	n2 := append([]int{}, nums[n/2:]...)
	// 递归排序后，这个数据已经是有序的
	cnt := reversePairs(n1) + reversePairs(n2)
	// 统计归并前，两个子数组满足规则的个数
	j := 0
	for _, v := range n1 {
		for j < len(n2) && v > 2*n2[j] {
			// j 代表索引位置，同时也代表个数
			j++
		}
		cnt += j
	}
	// 归并合并
	p1, p2 := 0, 0
	for i := range nums {
		// p1<len(n1) 避免 n1[p1]溢出
		//  p2==len(n2) 避免n2[p2] 溢出
		if p1 < len(n1) && (p2 == len(n2) || n1[p1] < n2[p2]) {
			nums[i] = n1[p1]
			p1++
		} else {
			nums[i] = n2[p2]
			p2++
		}
	}
	return cnt
}

func main() {
	fmt.Println(reversePairs([]int{1,3,2,3,1}))
}
