package main

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == ""{
		return 0
	}
	needleLen := len(needle)
	for i:=0;i<len(haystack);i++{
		if haystack[i] == needle[0] && i+needleLen <=len(haystack)&& haystack[i:i+needleLen] == needle{
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("helalolgll","ll"))
}
