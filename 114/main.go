package main

// TODO
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

var last *TreeNode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Right)
	flatten(root.Left)
	root.Right = last
	root.Left = nil
	last = root
}

type S struct {
	name string
}

type Result struct {
	Status int
}

//func main() {
//	timeout := 3 * time.Second
//	ctx, cancel := context.WithTimeout(context.Background(), timeout)
//	defer cancel()
//
//	select {
//	case <-time.After(1 * time.Second):
//		fmt.Println("waited for 1 sec")
//	case <-time.After(2 * time.Second):
//		fmt.Println("waited for 2 sec")
//	case <-time.After(3 * time.Second):
//		fmt.Println("waited for 3 sec")
//	case <-ctx.Done():
//		fmt.Println(ctx.Err())
//	}
//}




func main() {

	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))
}

func DeferFunc1(i int) (t int) {
	t = i
	defer func(t int) {
		t += 3
	}(t)
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}