// Реализация стека с помощью односвязанного списка
package main

import "fmt"

type Node struct {
	prev *Node
	Data int
}

type List struct {
	len  int
	head *Node
	tail *Node
}

func (l *List) Head() *Node {
	return l.head
}

func (l *List) Tail() *Node {
	return l.tail
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Find(elem int) *Node {
	if l.head != nil {
		for tmp := l.head; tmp != l.tail; tmp = tmp.prev {
			if tmp.Data == elem {
				return tmp
			}
		}
	}
	return nil
}

func (l *List) Add(node *Node) {
	l.len++
	if l.head == nil {
		l.head = node
		l.tail = l.head
		return
	}
	node.prev = l.head
	l.head = node
}

func (l *List) Delete(node *Node) {
	l.len--
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	}
	l.head = node.prev
}


type Stack struct {
	list *List
}

func NewStack() *Stack {
	return &Stack{
		list: &List{},
	}
}

func (s *Stack) Push(elem int) {
	node := &Node{
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
	//stack.Push(8)
	//stack.Push(9)
	//stack.Push(10)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
