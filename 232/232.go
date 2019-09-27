package main

import "fmt"

type Node struct {
	Value int
	Next *Node
}


type MyQueue struct {
	head *Node
	tail *Node
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	if this.tail == nil{
		this.tail = &Node{Value:x}
		this.head = this.tail
	}else {
		this.tail.Next = &Node{Value:x}
		this.tail = this.tail.Next
	}
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	value:=this.head.Value
	this.head = this.head.Next
	if this.head == nil{
		this.tail = nil
	}
	return value
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.head.Value
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.tail == nil
}

func main() {
	obj := Constructor()
	obj.Push(1)
	fmt.Println(obj.Peek())
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */