package main

import "fmt"

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	pre,cur := 1,1
	for i:=1;i<len(s);i++ {
		tmp:=cur
		if s[i] == '0' {
			if  s[i-1] == '1' || s[i-1] == '2'{
				cur = pre
			}else {
				return 0
			}
		}else if s[i-1] =='1' || s[i-1]=='2' && s[i] - '0' <=6 {
			cur = cur + pre
			// 如果当前的字符 可以和前一个组成合法字符，那么这里就有 组成合法字符的个数等于 dp[i-2]  单个字符的时候，字符个数等于dp[di-1]  于是当前的字符个数等于 dp[i] = dp[i-2] dp[i-1]
		}
		pre = tmp

	}
	return cur
}

func main() {
	fmt.Println(numDecodings("10110"))
}
