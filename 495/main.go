package main

import "fmt"

func findPoisonedDuration(timeSeries []int, duration int) int {
	var d int
	res := make([][]int, 0)
	intervals := make([][]int, len(timeSeries))
	for i := 0; i < len(timeSeries); i++ {
		intervals[i] = []int{timeSeries[i], timeSeries[i] + duration - 1}
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

	for i := 0; i < len(res); i++ {
		d += res[i][1] - res[i][0] + 1
	}
	return d
}

func main() {
	fmt.Println(findPoisonedDuration([]int{1,2,3,6,8,9,12,15},2))
}