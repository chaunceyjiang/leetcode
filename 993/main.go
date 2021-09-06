package main

import (
	"fmt"
	"github.com/chaunceyjiang/leetcode/utils"
)

type TreeNode = utils.TreeNode

type stack []*TreeNode

func (s *stack) Pop() *TreeNode {
	res := (*s)[0]
	*s = (*s)[1:]
	return res
}
func (s *stack) Push(node *TreeNode) {
	*s = append(*s, node)
}
func (s stack) Len() int {
	return len(s)
}

func newStack() stack {
	return make([]*TreeNode, 0)
}

var s = newStack()
var xDepth, yDepth = 0, 0
var xParent, yParent *TreeNode = nil, nil

func isCousins(root *TreeNode, x int, y int) bool {
	s.Push(root)
	 levelFind(0, x, y)
	return　i
}

func levelFind(depth int, x, y int) {
	count := 0
	curLen := s.Len()
	if curLen == 0 {
		return
	}
	for i := 0; i < curLen; i++ {
		curNode := s.Pop()
		if curNode.Val == x || curNode.Val == y {
			count++
		}
		if curNode.Left != nil {
			s.Push(curNode.Left)
		}
		if curNode.Right != nil {
			s.Push(curNode.Right)
		}
	}

	levelFind(depth+1, x, y)
}

func main() {
	tree := utils.BuildTree([]int{1, 2, 4, 3, 5}, []int{2, 4, 1, 3, 5})
	fmt.Println(isCousins(tree, 5, 4))
}
