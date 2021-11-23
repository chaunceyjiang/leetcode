package main

import "fmt"

func buddyStrings(s string, goal string) bool {
	// 字符串长度不同,false
	if len(s) != len(goal) {
		return false
	}
	m := make(map[int32]int)
	if s == goal {
		max := 0
		for _, x := range s {
			m[x]++
			if m[x] > max {
				max = m[x]
			}
		}
		if max >= 2 {
			return true
		}
		return false
	}

	chars := []byte(s)
	index := make([]int, 0)
	for i := 0; i < len(chars); i++ {
		if chars[i] != goal[i] {

			index = append(index, i)
		}
	}

	if len(index) != 2 {
		return false
	}
	chars[index[0]], chars[index[1]] = chars[index[1]], chars[index[0]]
	return string(chars) == goal
}

func main() {
	fmt.Println(buddyStrings("aaabc", "aabac"))
}
