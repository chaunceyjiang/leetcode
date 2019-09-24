package main

/*import "fmt"
// TODO
func insert(intervals [][]int, newInterval []int) [][]int {
	left := midFind(0, len(intervals)-1, newInterval[0], intervals, 0)
	right := midFind(0, len(intervals)-1, newInterval[1], intervals, 1)
	l := max(left-1, 0)
	r := min(right, len(intervals))
	res := make([][]int, 0)
	res = append(res, intervals[:l]...)
	midlis := intervals[l:r]
	if len(midlis) == 0 {
		res = append(res, newInterval)
	} else {
		fmt.Println(midlis)
		r:=make([]int,2)
		if newInterval[0] <= midlis[0][1] && newInterval[1] >=midlis[len(midlis)-1][0] {
			r[0] = min(midlis[0][0],newInterval[0])
			r[1] = max(midlis[len(midlis)-1][1],newInterval[1])
		}
		if
	}
	res = append(res, intervals[r:]...)
	return res
}
func midFind(l, r, k int, nums [][]int, i int) int {
	mid := (l + r) / 2
	if l >= r {
		return mid
	}
	if k == nums[mid][i] {
		return mid
	} else if k > nums[mid][i] {

		return midFind(mid+1, r, k, nums, i)
	} else {
		return midFind(l, mid, k, nums, i)
	}

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {
	var test = [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	fmt.Println(midFind(0, len(test)-1, 4, test, 0))
	fmt.Println(midFind(0, len(test)-1, 8, test, 1))
	fmt.Println(insert(test,[]int{4,8}))
}
*/