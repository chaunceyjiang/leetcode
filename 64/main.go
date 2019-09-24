package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func minPathSum(grid [][]int) int {
	dp := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				dp[i] = grid[i][j]
			} else if i == 0 && j !=0 {
				dp[j] = dp[j-1] + grid[i][j]
			} else if j == 0 && i != 0 {
				dp[j] = dp[j] + grid[i][j]
			} else {
				// dp[j-1]+grid[i][j] 表示从右向左走
				// dp[j]+grid[i][j] 表示向上走
				dp[j] = min(dp[j]+grid[i][j], dp[j-1]+grid[i][j])
			}
		}
	}

	return dp[len(dp)-1]
}

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
