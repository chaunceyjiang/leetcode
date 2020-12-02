package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的迭代遍历,利用一个中间数组和栈进行处理
// 前序输出
func beforePostorderTraversal(root *TreeNode) []int {
	stk := make([]*TreeNode, 0) // 栈
	res := make([]int, 0)
	p := root
	for p != nil || len(stk) > 0 {
		if p != nil {
			res = append(res, p.Val)
			stk = append(stk, p)
			p = p.Left
		} else {
			t:= stk[len(stk)-1]
			p = t.Right
			stk=stk[:len(stk)-1]
		}
	}
	return res
}

// 中序输出
func MidpostorderTraversal(root *TreeNode) []int {
	stk := make([]*TreeNode, 0) // 栈
	res := make([]int, 0)
	p := root
	for p != nil || len(stk) > 0 {
		if p != nil {
			stk = append(stk, p)
			p = p.Left
		} else {
			t:= stk[len(stk)-1]
			res = append(res, t.Val)
			p = t.Right
			stk=stk[:len(stk)-1]
		}
	}
	return res
}

// 前序-->插入头部，然后从右边开始-->后序输出
func AfterpostorderTraversal(root *TreeNode) []int {
	stk := make([]*TreeNode, 0) // 栈
	res := make([]int, 0)
	p := root
	for p != nil || len(stk) > 0 {
		if p != nil {
			stk = append(stk, p)
			// 插入头部，然后从右边开始
			res = append([]int{p.Val}, res...)
			p = p.Right
		} else {
			t:=stk[len(stk)-1]
			p = t.Left
			stk=stk[:len(stk)-1]
		}
	}
	return res
}
func main() {
	fmt.Println([]int{1}[1:])
}
