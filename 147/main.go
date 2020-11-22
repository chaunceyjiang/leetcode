package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := new(ListNode)
	dummyHead.Next = head
	cur := head.Next
	lastSort := head
	for cur != nil {
		if lastSort.Val <= cur.Val {
			lastSort = lastSort.Next
		} else {
			p := dummyHead
			for p.Next.Val <= cur.Val {
				// 找到比当前大的值，插入到它的前面
				p = p.Next
			}
			// 排序下一个是当前的下一个
			lastSort.Next = cur.Next

			// 交换
			//p.Next.Next = cur.Next
			cur.Next=p.Next
			p.Next=cur
		}
		cur = lastSort.Next
	}
	return dummyHead.Next
}

func createList(nums []int) *ListNode {
	var head, p *ListNode

	for i := 0; i < len(nums); i++ {
		t := new(ListNode)
		t.Val = nums[i]
		if p == nil {
			p = t
			head = p
		} else {
			p.Next = t
			p = p.Next
		}
	}
	return head
}

func printList(root *ListNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	printList(root.Next)
}
func main() {
	l := createList([]int{4,2,1,3})

	printList(insertionSortList(l))
}
