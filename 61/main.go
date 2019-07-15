package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil{
		return nil
	}
	p := head
	var tail = head
	l := 0
	for p != nil {
		l++
		tail = p
		p = p.Next

	}
	tail.Next = head
	k %= l
	k = l -k
	for i:=0;i<k;i++{
		head = head.Next
		tail = tail.Next
	}
	tail.Next = nil
	return head
}
