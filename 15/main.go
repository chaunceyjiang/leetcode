package main

import "fmt"

//func threeSum(nums []int) [][]int {
//	if len(nums)<3{
//		return nil
//	}
//	QuickSort(0,len(nums)-1,nums)
//
//	for i := 0; i < len(nums);i++ {
//
//	}
//}

func QuickSort(left,right int,nums []int) {
	i, j := left, right
	t:=nums[left]
	for i < j {

		for i <j && nums[j] > t{
			j--
		}
		for i <j && nums[i] <= t{
			i++
		}
		swap(i,j,nums)
	}
	swap(left,i,nums)
	QuickSort(left,i-1,nums)
	QuickSort(i+1,right,nums)
}
func swap(x, y int, nums []int) {
	t := nums[x]
	nums[x] = nums[y]
	nums[y] = t
}

func main() {
	s:="123244"

	fmt.Printf("%T",s[0] )
}
