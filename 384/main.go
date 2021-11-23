package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Solution struct {
	nums   []int
	result []int
	length int
}

func Constructor(nums []int) Solution {
	s:=Solution{nums: nums, result: make([]int, len(nums))}
	s.Reset()
	return s

}

func (this *Solution) Reset() []int {
	this.length = len(this.nums)
	this.result = make([]int, len(this.nums))
	copy(this.result, this.nums)
	return this.nums
}

func (this *Solution) Shuffle() []int {
	rand.Seed(time.Now().UnixNano())
	 rand.Shuffle(len(this.result), func(i, j int) {
		this.result[i],this.result[j] = this.result[j],this.result[i]
	})
	return this.result
}

func main() {
	c := Constructor([]int{1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Println(c.Shuffle())
	fmt.Println(c.Shuffle())
	fmt.Println(c.Shuffle())
}
