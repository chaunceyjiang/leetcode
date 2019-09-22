package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func buildTree(preorder []int, inorder []int) *TreeNode {
	index:=0
	return createTree(preorder,inorder,0,len(inorder)-1,&index)
}

func createTree(preorder []int, inorder []int, start int, end int,index *int) *TreeNode {
	if start > end{
		return nil
	}
	rootVal:=preorder[*index]
	*index = *index +1
	for i:=start;i<=end;i++{
		if inorder[i] == rootVal{
			root:=&TreeNode{Val:rootVal}
			root.Left = createTree(preorder,inorder,start,i-1,index)
			root.Right = createTree(preorder,inorder,i+1,end,index)
			return root
		}
	}
	return nil
}
