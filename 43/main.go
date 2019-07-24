package main

import "fmt"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	res := make([]uint8, len(num1)+len(num2)+1)
	for i,t1 := len(num1)-1,0; i>=0; i,t1=i-1,t1+1 {
		for j :=len(num1)-1; j>=0; j=j-1 {
			r,z:=mut(num1[i], num2[j], res[t1])
			res[t1+1] = r
			res[t1] = z
		}
	}
	fmt.Println(res)
	return ""

}

//func () {
//
//}
func mut(x, y, z uint8) (uint8, uint8) {
	x-=48
	y-=48
	m := x*y + z
	return m / 10, m % 10
}

func main() {
	fmt.Println(multiply("92", "87"))
	fmt.Println(multiply("1213142435", "323823293729329372"))
}
