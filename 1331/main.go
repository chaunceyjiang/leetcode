package main

import (
	"fmt"
	"sort"
)

type ByValue []int

func (b ByValue) Len() int {
	return len(b)
}

func (b ByValue) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b ByValue) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func arrayRankTransform(arr []int) []int {
	m := make(map[int][]int)
	arrNew := make([]int, 0)
	for i, v := range arr {
		if _, ok := m[v]; !ok {
			m[v] = []int{i}
			arrNew = append(arrNew, v)
		} else {
			m[v] = append(m[v], i)
		}
	}
	sort.Sort(ByValue(arrNew))
	for i, v := range arrNew {
		for _, index := range m[v] {
			arr[index] = i+1
		}
	}
	return arr
}
func arrayRankTransform2(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return []int{}
	}
	index := make([]int, n)
	ans := make([]int, n)
	for i := range index {
		index[i] = i
	}

	sort.Slice(index, func(n1, n2 int) bool { return arr[index[n1]] < arr[index[n2]] })

	ans[index[0]] = 1
	for i := 1; i < len(index); i++ {
		if arr[index[i]] == arr[index[i-1]] {
			ans[index[i]] = ans[index[i-1]]
		} else {
			ans[index[i]] = ans[index[i-1]] + 1
		}
	}
	return ans
}
func main() {
	fmt.Println(arrayRankTransform2([]int{1,1,1}))
}
