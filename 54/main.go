package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	res := make([]int, 0)
	var d = 0
	var x, y = -1, -1
	var max = 0
	x_n := len(matrix)
	y_n := len(matrix[0])
	flag := make([][]bool, x_n)
	for i := 0; i < x_n; i++ {
		flag[i] = make([]bool, y_n)
	}
	for max < x_n*y_n {
		if d == 0 {
			x++
			y++
			for y < y_n && !flag[x][y] {
				res = append(res, matrix[x][y])
				flag[x][y] = true
				y++
				max++
			}
			d = 1
		}
		if d == 1 {
			x++
			y--
			for x < x_n && !flag[x][y] {
				res = append(res, matrix[x][y])
				flag[x][y] = true
				x++
				max++
			}
			d = 2
		}
		if d == 2 {
			x--
			y--
			for y >= 0 && !flag[x][y] {
				res = append(res, matrix[x][y])
				flag[x][y] = true
				y--
				max++
			}

			d = 3
		}
		if d == 3 {
			y++
			x--
			for x >= 0 && !flag[x][y] {
				res = append(res, matrix[x][y])
				flag[x][y] = true
				x--
				max++
			}
			d = 0
		}
	}

	return res
}

func main() {
	res := make([][]int, 3)
	res[0] = []int{1, 2, 3, 4}
	res[1] = []int{5, 6, 7, 8}
	res[2] = []int{9, 10, 11, 12}
	fmt.Println(spiralOrder(res))
}

// 剑指 Offer 54. 二叉搜索树的第k大节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	var i = 0
	val := searchNone(root, &i, k)
	return val.Val
}

func searchNone(root *TreeNode, i *int, k int) *TreeNode {
	if root == nil {
		return nil
	}
	var val1,val2 *TreeNode
	val1 = searchNone(root.Right, i, k)
	*i += 1
	if *i == k {
		return root
	}
	val2 = searchNone(root.Left, i, k)
	if val1 != nil{
		return val1
	}
	return val2
}
