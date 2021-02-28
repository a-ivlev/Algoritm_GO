// Реализация очереди с помощью односвязанного списка
package main

import "fmt"

type Node struct {
	next *Node
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
		for tmp := l.tail; tmp != l.head; tmp = tmp.next {
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
	l.head.next = node
	l.head = node
}

func (l *List) Delete(node *Node) {
	l.len--
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	}
		l.tail = node.next
}

type Queue struct {
	list *List
}

func NewQueue() *Queue {
	return &Queue{
		list: &List{},
	}
}

func (s *Queue) Push(elem int) {
	node := &Node{
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
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
}
