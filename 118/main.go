package main

import "fmt"

func generate(numRows int) [][]int {
	res := make([][]int,0)
	res = append(res, []int{1})

	if numRows == 1 {
		return res
	}
	res = append(res, []int{1, 1})

	if numRows == 2 {
		return res
	}
	for i := 3; i <= numRows; i++ {
		temp := make([]int, i)
		for j := 0; j < i; j++ {
			if j == 0 || j == i-1 {
				temp[j] = 1
			} else {
				temp[j] = res[i-2][j] + res[i-2][j-1]
			}
		}
		res = append(res, temp)
	}
	return res
}

func main() {
	fmt.Println(generate(7))
}
