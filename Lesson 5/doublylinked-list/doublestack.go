package main

import (
	"fmt"
	l "Algoritm_GO/lesson-5/doublelinked"
)

type Stack struct {
	list *l.List
}

func NewStack() *Stack {
	return &Stack{
		list: &l.List{},
	}
}

func (s *Stack) Push(elem int) {
	node := &l.Node{
		Data: elem,
	}
	s.list.Add(node)
}
func (s *Stack) Pop() int {
	if s.list.Len() == 0 {
		return 0
	}
	elem := s.list.Head().Data
	s.list.Delete(s.list.Head())
	return elem
}

func main() {
	stack := NewStack()
	stack.Push(5)
	stack.Push(6)
	stack.Push(7)
	stack.Push(8)
	stack.Push(9)
	stack.Push(10)
	fmt.Println("POP", stack.Pop())
	fmt.Println("POP", stack.Pop())
	//fmt.Println("POP", stack.Pop())
	fmt.Println(stack.list.Head())
	fmt.Println(stack.list.Tail())
	fmt.Println(stack.list.Len())
}

