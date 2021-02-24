package main

import (
	"fmt"
	l "Algoritm_GO/lesson-5/doublelinked"
)

type Queue struct {
	list *l.List
}

func NewQueue() *Queue {
	return &Queue{
		list: &l.List{},
	}
}

func (s *Queue) Push(elem int) {
	node := &l.Node{
		Data: elem,
	}
	s.list.Add(node)
}
func (s *Queue) Pop() int {
	if s.list.Len() == 0 {
		return 0
	}
	elem := s.list.Tail().Data
	s.list.Delete(s.list.Tail())
	return elem
}

func main() {
	queue := NewQueue()
	queue.Push(5)
	queue.Push(6)
	queue.Push(7)
	queue.Push(8)
	queue.Push(9)
	queue.Push(10)
	//fmt.Println(queue.list.Len())
	//fmt.Println(queue.list.Head())
	//fmt.Println(queue.list.Tail())
	fmt.Println("POP", queue.Pop())
	fmt.Println("POP", queue.Pop())
	fmt.Println("POP", queue.Pop())
	fmt.Println(queue.list.Head())
	fmt.Println(queue.list.Tail())
	fmt.Println(queue.list.Len())
}