package doublelinked

type Node struct {
	prev *Node
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

func (l *List) FindHead(elem int) *Node {
	if l.head != nil {
		for tmp := l.head; tmp != l.tail; tmp = tmp.prev {
			if tmp.Data == elem {
				return tmp
			}
		}
	}
	return nil
}

func (l *List) Find(elem int) *Node {
	if l.tail != nil {
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
			l.tail = node
			return
		}
		node.prev = l.head
		l.head.next = node
		l.head = node
		return
}

func (l *List) Delete(node *Node) {
	l.len--
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	}
	if node == l.head {
		l.head = node.prev
	}
	if node == l.tail {
		l.tail = node.next
	}
}