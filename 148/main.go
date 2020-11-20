package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归合并
func sortList(head *ListNode) *ListNode {
	var fast, slow, t *ListNode
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	fast = head
	slow = head
	t = head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		t = slow
		slow = slow.Next

	}
	t.Next = nil

	return migre(sortList(head), sortList(slow))
}

func migre(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head *ListNode
	if l1.Val < l2.Val {
		head = l1
		l1.Next = migre(l1.Next, l2)
	} else {
		head = l2
		l2.Next = migre(l1, l2.Next)
	}

	return head

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

/* 归并排序
step=1: (3->4)->(1->7)->(8->9)->(2->11)->(5->6)
step=2: (1->3->4->7)->(2->8->9->11)->(5->6)
step=4: (1->2->3->4->7->8->9->11)->(5->6)
step=8: (1->2->3->4->5->6->7->8->9->11)
*/
func sortList2(root *ListNode) *ListNode {
	head := new(ListNode)
	length := 0
	head.Next = root
	p := root
	// 计算长度
	for p != nil {
		p = p.Next
		length++
	}
	// 切分链表
	for i := 1; i < length; i *= 2 {
		cur := head.Next
		tail := head
		for cur != nil {
			// 切分为 单个节点
			left := cur
			right := cut(left, i)
			//下一个开始
			cur =cut(right,i)
			// 将单个节点合并
			tail.Next = megre2(left, right)
			for tail.Next != nil {
				tail = tail.Next
			}
		}
	}
	return head.Next

}

// megre2 双路合并
func megre2(l1, l2 *ListNode) *ListNode {
	p := new(ListNode)
	head := p
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
			p = p.Next
		} else {
			p.Next = l2
			l2 = l2.Next
			p = p.Next
		}
	}
	if l1 == nil {
		p.Next = l2
	} else {
		p.Next = l1
	}
	return head.Next
}

// cut 切分 链表
func cut(root *ListNode, n int) *ListNode {
	p := root
	for p != nil && n > 1 {
		n--
		p = p.Next
	}
	// 如果p 是nil 则表示后面没有其他节点
	if p == nil {
		return nil
	}
	next := p.Next
	p.Next = nil
	return next

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
	t := sortList2(l)
	printList(t)
}
