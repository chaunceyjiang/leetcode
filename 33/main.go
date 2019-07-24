package main

import "fmt"

func search(nums []int, target int) int {
	min := nums[0]
	mid := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
			mid = i
		}
	}
	mid = midSearch(0,len(nums)-1,nums,target)
	return mid
}

func midSearch(l, r int, nums []int,target int) int{
	if l ==r {
		if nums[l] == target{
			return l
		}else {
			return -1
		}

	}
	mid :=(l+r)/2
	if target<= nums[mid]{
		return midSearch(l,mid,nums,target)
	} else {
		return midSearch(mid,r,nums,target)
	}

}
func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}
