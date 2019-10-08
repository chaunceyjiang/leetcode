package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[0]); j++ {
			if i == 0 && j == 0 {
				if obstacleGrid[i][j] == 0 {
					obstacleGrid[i][j] = 1
				} else {
					obstacleGrid[i][j] = 0
				}
				continue
			}

			if i == 0 {
				if obstacleGrid[i][j] == 1 {
					obstacleGrid[i][j] = 0
				} else {
					obstacleGrid[i][j] = obstacleGrid[i][j-1]
				}

			}
			if j == 0 {
				if obstacleGrid[i][j] == 1 {
					obstacleGrid[i][j] = 0
				} else {
					obstacleGrid[i][j] = obstacleGrid[i-1][j]
				}
			}
			if i != 0 && j != 0 {
				if obstacleGrid[i][j] == 1 {
					obstacleGrid[i][j] = 0
				} else {
					// 在原表的基础上进行计数
					// 原理： 当前位置为 1（障碍） 则把当前位置制为0 表示当前没有路可以到达这里
					// 当前位置为 0（可到达） 则将 上面 和左边的条数相加。
					obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
				}
			}
		}
	}
	return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
