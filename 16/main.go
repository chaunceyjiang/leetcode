package main

import "fmt"

func threeSumClosest(nums []int, target int) int {
	if len(nums) <3{
		return 0
	}

	quickSort(0,len(nums)-1,nums)
	var res  = nums[0] + nums[1] + nums[2]
	for i:=0;i<len(nums)-2;i++{
		head := i+1
		tail :=len(nums)-1
		for head<tail{
			sum:=nums[i] + nums[head] + nums[tail]
			if abs(sum - target) < abs( res - target){
				res = sum
			}
			if sum > target{
				tail--
			}
			if sum < target {
				head++
			}
			if sum == target{
				return sum
			}
		}
	}
	return res
}
func abs(a int) int {
	if a< 0 {
		return -a
	}
	if a == 0{
		return 0
	}
	return a
}
func quickSort(l, r int, nums []int) {
	if l>=r{
		return
	}
	t:=nums[l]
	i,j:=l,r
	for i<j  {

		for i < j && nums[j] > t {
			j--
		}
		for i < j && nums[i] <= t{
			i++
		}
		swap(i,j,nums)
	}
	swap(i,l,nums)
	quickSort(l,i-1,nums)
	quickSort(i+1,r,nums)

}

func swap(x, y int, nums []int) {
	t:=nums[x]
	nums[x] = nums[y]
	nums[y] = t
}

func main() {
	fmt.Println(threeSumClosest([]int{1,5,0,7,2,3,5},19))
}