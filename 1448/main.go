package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	var count = 0
	dfs(root, root.Val, &count)
	return count
}
func dfs(root *TreeNode, pervVal int, count *int) {
	if root == nil {
		return
	}
	if root.Val >= pervVal {
		*count += 1
		pervVal = root.Val
	}
	dfs(root.Left, pervVal, count)
	dfs(root.Right, pervVal, count)
}

func main() {

}
