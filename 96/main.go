package main

func numTrees(n int) int {
	/*
	   卡塔兰数
	G(n) = G(0)G(n-1) + G(1)G(n-2) + ... + G(n-2)G(1) + G(n-1)G(0)
	*/

	c := 1
	for i := 0; i < n; i++ {
		c = c * 2 * (2*i + 1) / (i + 2)
	}
	return c

}
