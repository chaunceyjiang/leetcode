package main

import (
	"fmt"
	"strings"
)

type Node struct {
	Value int
	Prev  *Node
}
type Stack struct {
	tail *Node
}

func (s *Stack) Push(value int) {
	if s.tail == nil {
		s.tail = &Node{Value: value}
	} else {
		s.tail = &Node{Value: value, Prev: s.tail}
	}
}
func (s *Stack) Pop() int {
	value := s.tail.Value
	s.tail = s.tail.Prev
	return value
}
func reverseParentheses(s string) string {
	var p = &Stack{}

	ss := strings.Split(s, "")
	var res = make([]string, 0)
	for i, v := range ss {
		switch v {
		case "(":
			p.Push(i)
		case ")":
			start := p.Pop()
			end := i
			reverse(ss, start, end)
		}
	}

	//reverse(ss,0,len(ss)-1)
	for _, v := range ss {
		if v == "(" || v == ")" {
			continue
		}
		res = append(res,v)
	}
	return strings.Join(res, "")
}

func reverse(s []string, start, end int) {
	start++
	end--
	for start < end {
		t := s[start]
		s[start] = s[end]
		s[end] = t
		start++
		end--
	}

}

func main() {
	fmt.Println(reverseParentheses("a(bcdefghijkl(mno)p)q"))
}
