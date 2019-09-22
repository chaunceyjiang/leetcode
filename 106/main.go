package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func buildTree(inorder []int, postorder []int) *TreeNode {
	 index := len(postorder)-1
	 return createTree(inorder,postorder,0,index,&index) // 中序的位置
}

func createTree(inorder []int, postorder []int,start,end int,index *int ) *TreeNode {
	if start > end {
		return nil
	}
	rootVal:=postorder[*index]
	*index = *index -1
	for i:=start;i<=end;i++{
		if inorder[i] == rootVal{ // 找到根节点在中序的位置，然后讲中序数组分成两半
			root := &TreeNode{Val:rootVal}
			// 优先构造右边，因为依赖全局 index 查找根节点的位置
			root.Right = createTree(inorder,postorder,i+1,end,index) //右边一半
			root.Left = createTree(inorder,postorder,start,i-1,index) //左边一半

			return root
		}
	}
	return nil
}
func main() {
	root:=buildTree([]int{9,3,15,20,7},[]int{9,15,7,20,3})
	fmt.Println(root)
}