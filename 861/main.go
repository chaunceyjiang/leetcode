package main

import (
	"fmt"
	"math"
)

// 为了保证每一行的数最大，那么第一列应该全为1
func matrixScore(A [][]int) int {
	N := len(A)
	M := len(A[0])
	// 保证第一列 全是1
	res := float64(0)
	for i := 0; i < N; i++ {
		if A[i][0] == 0 {
			for j := 0; j < M; j++ {
				A[i][j] ^= 1
			}
		}
		res += math.Pow(2, float64(M-1))
	}
	// 保证第二列 1 比0多
	for j := 1; j < M; j++ {
		count := 0
		for i := 0; i < N; i++ {
			count += A[i][j]
		}
		// 1 的个数小于 0 的个数
		for i := 0; i < N; i++ {
			if count < (N)/2 {
				A[i][j] ^= 1
			}
			if A[i][j] == 1 {
				res += math.Pow(2, float64(M-1-j))
			}
		}

	}

	return int(res)
}

func main() {
	A:=make([][]int,4)
	A[0] = []int{0,1}
	A[1] = []int{0,1}
	A[2] = []int{0,1}
	A[3] = []int{0,0}
	fmt.Println(matrixScore(A))
}
