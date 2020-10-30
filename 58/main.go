package main

import "fmt"

func lengthOfLastWord(s string) int {
	sLen := len(s)
	if sLen == 0{
		return 0
	}
	endWord := sLen
	for i := sLen; i > 0; i-- {
		if s[i-1:i] != " " {
			endWord = i
			break
		}
	}
	startWord := 0
	for i := endWord - 1; i >= 0; i-- {
		if s[i:i+1] == " " {
			startWord = i + 1
			break
		}
	}
	return endWord - startWord
}

func main() {
	fmt.Println(lengthOfLastWord(" ad "))
}
