package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type MyLinkedList struct {
	count int
	head  *Node
	tail  *Node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.count {
		return -1
	}

	// 这里性能较低
	for p, i := this.head, 0; p != nil; p, i = p.Next, i+1 {
		if i == index {
			return p.Value
		}
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	var node = &Node{
		Value: val,
		Next:  this.head,
	}

	this.count++

	if this.head == nil {
		this.tail = node
	}

	this.head = node
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	var node = &Node{
		Value: val,
		Next:  nil,
	}

	this.count++
	if this.head == nil {
		this.head = node
		this.tail = this.head
	} else {
		this.tail.Next = node
		this.tail = node
	}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.count {
		return
	}

	if index == this.count {
		this.AddAtTail(val)
		return
	}
	if index <= 0 {
		this.AddAtHead(val)
		return
	}

	var pp, p = this.get(index)
	if p != nil {
		this.count++
		var node = &Node{
			Value: val,
			Next:  p,
		}
		if pp == nil {
			this.head = node
		} else {
			pp.Next = node
		}

	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	var pp, p = this.get(index)
	if p != nil {
		this.count--

		if pp == nil {
			this.head = p.Next
		} else {
			pp.Next = p.Next
			if pp.Next == nil{
				this.tail = pp
			}
		}
	}
}
func (this *MyLinkedList) get(index int) (*Node, *Node) {

	var pp *Node
	for p, i := this.head, 0; p != nil; p, i = p.Next, i+1 {
		if i == index {
			return pp, p
		}
		pp = p
	}
	return nil, nil
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
func main() {
	obj := Constructor()
	obj.AddAtHead(8)
	fmt.Println(obj.Get(1))
	obj.AddAtTail(81)
	obj.DeleteAtIndex(2)
	obj.AddAtHead(26)
	obj.DeleteAtIndex(2)
	fmt.Println(obj.Get(1))
	obj.AddAtTail(24)
	obj.AddAtHead(15)
	obj.AddAtTail(0)
	obj.AddAtTail(13)
	obj.AddAtTail(1)
	obj.AddAtIndex(6,33)
	fmt.Println(obj.Get(6))

}
