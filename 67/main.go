package main

import "fmt"

func addBinary(a string, b string) string {
	var s, x, c = "", "", "0"
	aLen := len(a) - 1
	bLen := len(b) - 1
	if aLen > bLen {
		b = fill(b, aLen-bLen)
	}
	if aLen < bLen {
		a = fill(a, bLen-aLen)
		aLen = bLen
	}
	for i := aLen; i >= 0; i-- {
		c, x = sum(a[i:i+1], b[i:i+1], c)
		s += x
	}
	if c == "1" {
		s = s + "1"
	}
	return reverse(s)
}
func fill(s string, x int) string {
	var t = ""
	for i := 0; i < x; i++ {
		t += "0"
	}
	return t + s
}
func sum(a, b, c string) (string, string) {
	if a == "1" && b == "1" {
		return "1", c
	}
	if a == "0" && b == "0" {
		return "0", c
	}
	if c == "0" {
		return "0", "1"
	}
	return "1", "0"
}
func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
func main() {
	fmt.Println(addBinary("11100101000101010111110010010101010",
		"11000010101010101010100101010101010101010101010101010101010100000000000111111111111111111111111111111000000111111100000000111111111111"))
}
