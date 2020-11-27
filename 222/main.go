package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	var t int
	t = 0
	count(root, &t)
	return t
}
func count(root *TreeNode, t *int) {
	if root == nil {
		return
	}
	*t++
	count(root.Left, t)
	count(root.Right, t)
}
func main() {

}
