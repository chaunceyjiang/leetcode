package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	index := make([]uint8, 26)
	for i := 0; i < len(s); i++ {
		index[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		index[t[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if index[i] != 0 {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println('a'&'b'&'c' == 'y'&'b'&'a')
	fmt.Println(isAnagram("abc", "ba"))
	fmt.Println("a"[0] - 'a')
}
