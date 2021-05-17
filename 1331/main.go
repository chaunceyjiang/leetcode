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
	sort.Slice()
	sort.Sort(ByValue(arrNew))
	for i, v := range arrNew {
		for _, index := range m[v] {
			arr[index] = i+1
		}
	}
	return arr
}
func main() {
	fmt.Println(arrayRankTransform([]int{40, 10, 20, 30}))
}
