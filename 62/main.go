package main

import "fmt"

func uniquePaths(m int, n int) int {
	dp :=make([][]int,m)
	for i:=0;i<m;i++{
		dp[i] = make([]int,n)
	}
	// 从左往右，从上到下扫描
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if i == 0 || j==0{
				// 边界都是 1
				dp[i][j] = 1
			}else{
				// 因为只能从上到下，从左到右
				// 因此当前位置，只能是两个方向的总和
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

func main() {
	fmt.Println(uniquePaths(3,7))
}