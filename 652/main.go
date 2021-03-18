package main

import (
	"fmt"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	m := make(map[string]int)
	res := make([]*TreeNode, 0)
	dfs(root, m, &res)
	return res
}

func dfs(root *TreeNode, m map[string]int, res *[]*TreeNode) string {
	if root == nil {
		return "#"
	}
	cur := strings.Join([]string{fmt.Sprintf("%d", root.Val), dfs(root.Left, m, res), dfs(root.Right, m, res)}, ",")
	if m[cur] == 1 {
		*res = append(*res, root)
	}
	m[cur]++
	return cur
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	index := 0
	return createTree(preorder, inorder, 0, len(inorder)-1, &index)
}

func createTree(preorder []int, inorder []int, start int, end int, index *int) *TreeNode {
	if start > end {
		return nil
	}
	rootVal := preorder[*index]
	*index = *index + 1
	for i := start; i <= end; i++ {
		if inorder[i] == rootVal {
			root := &TreeNode{Val: rootVal}
			root.Left = createTree(preorder, inorder, start, i-1, index)
			root.Right = createTree(preorder, inorder, i+1, end, index)
			return root
		}
	}
	return nil
}

func main() {
	root := buildTree([]int{1, 2, 4, 3, 2, 4, 4}, []int{4, 2, 1, 4, 2, 3, 4})
	fmt.Println(findDuplicateSubtrees(root))
}
