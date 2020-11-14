package main

import "fmt"

func sortArrayByParityII(A []int) []int {
	for i,j := 0,1; i < len(A); i+=2 {
		if A[i]%2==1{
			// 当前为奇数
			for A[j]%2 != 0 {
				j+=2
			}
			A[i],A[j]=A[j],A[i]
		}
	}
	return A
}

func main() {
	fmt.Println(sortArrayByParityII([]int{4,2,5,7,1,2,4,5,66,55,46,23,13,15,17,24,26,28}))
}
