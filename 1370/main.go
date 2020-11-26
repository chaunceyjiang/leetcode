package main

import "fmt"

func sortString(s string) string {
	if len(s) <= 1 {
		return s
	}
	sMap := make([]int, 26)
	for i := 0; i < len(s); i++ {
		sMap[s[i]-'a']++
	}

	res := make([]byte, 0, len(s))
	for len(res) < len(s) {
		for i := 0; i < 26; i++ {
			if sMap[i] > 0 {
				sMap[i]--
				res = append(res, uint8(i)+'a')
			}
		}
		for i := 25; i >=0; i-- {
			if sMap[i] > 0 {
				sMap[i]--
				res = append(res, uint8(i)+'a')
			}
		}
	}
	return string(res)
}

func main() {
	fmt.Println(sortString("spo"))
}
