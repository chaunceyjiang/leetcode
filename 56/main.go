package main

import "fmt"

func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	for i := 0; i<len(intervals); i++ {
		for j := len(intervals)-1; j>=i; j-- {
			if intervals[i][0] > intervals[j][0] {
				swap(i, j, intervals)
			}

		}
		if len(res) != 0 {
			if intervals[i][1] <= res[len(res)-1][1] {
				continue
			}
			if intervals[i][0] <= res[len(res)-1][1] {
				res[len(res)-1][1] = intervals[i][1]
			} else {
				res = append(res, intervals[i])
			}
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}

func swap(x, y int, intervals [][]int) {
	t := intervals[x]
	intervals[x] = intervals[y]
	intervals[y] = t
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {8, 10}, {15, 18}, {2, 6}}))
}
