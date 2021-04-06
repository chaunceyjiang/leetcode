package main

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	flag := nums[0]
	count := 0
	index := 0
	replace:=flag
	for i := 0; i < len(nums); i++ {
		if nums[i] == flag {
			count += 1
			if count <= 2 {
				nums[index] = replace
				index +=1

			}
			continue
		}
		flag = nums[i]
		count = 1
		replace = nums[i]
		nums[index] = replace
		index +=1
	}

	return index
}

func removeDuplicates2(nums []int) int {
	var index int
	for _,val :=range nums {
		if index < 2 || val != nums[index-2] {
			nums[index]=val
			index++
		}
	}
	return index
}

func main() {

}
