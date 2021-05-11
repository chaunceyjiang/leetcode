package main

import "fmt"

/*
	A,B,C,D,E
 X = A^B^C^D^E  已知
 Y = B^C ^ D^E 奇数

A = X^Y

*/
func decode(encoded []int) []int {
	perm := make([]int, len(encoded)+1)
	Y, X := 0, 0
	for i := 1; i <= len(encoded)+1; i++ {
		Y ^= i
		if (i-1)%2 == 1 {
			X ^= encoded[i-1]
		}
	}
	perm[0] = X ^ Y
	for i := 1; i < len(perm); i++ {
		perm[i] = perm[i-1] ^ encoded[i-1]
	}
	return perm
}

func main() {
	fmt.Println(decode([]int{6, 5, 4, 6}))
}
