package main

import (
	"fmt"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	quickSort(0, len(nums)-1, nums)
	result := ""
	for _, v := range nums {
		result += strconv.Itoa(v)
	}
	if strings.TrimLeft(result,"0") == ""{
		result = "0"
	}
	return result
}

func gt(x, y int) bool {
	strX := strconv.Itoa(x)
	strY := strconv.Itoa(y)
	intX, _ := strconv.Atoi(strX + strY)
	intY, _ := strconv.Atoi(strY + strX)
	return intY > intX // 为真就交换
}

func le(x, y int) bool {
	strX := strconv.Itoa(x)
	strY := strconv.Itoa(y)
	intX, _ := strconv.Atoi(strX + strY)
	intY, _ := strconv.Atoi(strY + strX)
	return intY <= intX // 为真就交换
}

func quickSort(l, r int, nums []int) {
	if l >= r {
		return
	}
	t := nums[l]
	i, j := l, r
	for i < j {
		for i < j && gt(nums[j], t) {
			j--
		}

		for i < j && le(nums[i], t) {
			i++
		}

		swap(i, j, nums)
	}
	swap(l, i, nums)
	quickSort(l, i-1, nums)
	quickSort(i+1, r, nums)
}
func swap(x, y int, nums []int) {
	t := nums[x]
	nums[x] = nums[y]
	nums[y] = t
}
func main() {
	fmt.Println(largestNumber([]int{3,30,34,5,9}))
	fmt.Println(largestNumber([]int{0,0}))
}
