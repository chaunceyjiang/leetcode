package main

import "fmt"

func moveZeroes(nums []int)  {
	i,k:=0,0
	for ;i<len(nums);i++{
		if nums[i]!=0{
			nums[k]=nums[i]
			k++
		}
	}
	for ;k<len(nums);k++{
		nums[k]=0
	}
}
func main() {
	t:=[]int{1,2,3}
	moveZeroes(t)
	fmt.Println(t)
}
