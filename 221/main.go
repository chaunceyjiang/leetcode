package main

import "fmt"

func maximalSquare(matrix [][]byte) int {
	// dp[i][j] = min(dp[i-1][j-1],dp[i-1][j],dp[i][j-1])+1
	dp := make([][]int, len(matrix)+1)
	for i := 0; i <= len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0])+1)
	}

	max := 0
	for i := 1; i <= len(matrix); i++ {
		for j := 1; j <= len(matrix[0]); j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
				if dp[i][j] > max {
					max = dp[i][j]
				}
			}

		}
	}
	return max * max
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(maximalSquare([][]byte{
		{1, 0, 1, 0, 0},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 1, 0}}))
	//fmt.Println(maximalSquare([][]byte{
	//	{0,1},
	//	{1,0},
	//}))
}
