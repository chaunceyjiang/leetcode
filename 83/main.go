package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	pp := head
	p := head.Next
	for p != nil {
		if p.Val == pp.Val {
			pp.Next = p.Next
		}else{
			pp = pp.Next
		}
		p = p.Next
	}

	return head
}
func main() {

}
