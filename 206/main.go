package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pp *ListNode = nil
	var p = head

	for p != nil {
		t:=p
		p = p.Next
		t.Next = pp
		pp = t
	}
	return pp
}
