package main

import "fmt"

type Slice []int

func (s *Slice) Remove(value interface{}) {
	for i, v := range *s {
		if isEqual(value,v) {
			*s = append((*s)[:i],(*s)[i+1:]...)
		}
	}
}

func isEqual(value interface{}, v int) bool {
	vv := value.(int)
	return vv == v
}
func main() {
	s:=Slice([]int{1,2,3})
	s.Remove(2)
	fmt.Println(s)
}
