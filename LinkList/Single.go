package linklist

type Node[T any] struct {
	next  *Node[T]
	Value T
}
type LinkedList[T any] struct {
	root *Node[T]
	tail *Node[T]

	count int
}

// add node back
func (l *LinkedList[T]) PushBack(value T) {
	node := &Node[T]{
		Value: value,
	}

	l.count++

	if l.root == nil {
		l.root = node
		l.tail = node
		return
	}

	l.tail.next = node
	l.tail = node
}

func (l *LinkedList[T]) PushFront(value T) {
	node := &Node[T]{
		Value: value,
	}
	l.count++

	if l.root == nil {
		l.root = node
		l.tail = node
		return
	}
	node.next = l.root
	l.root = node
}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.root
}

func (l *LinkedList[T]) Back() *Node[T] {
	return l.tail
}

// O(N)
func (l *LinkedList[T]) Count() int {
	node := l.root
	cnt := 0

	for ; node != nil; node = node.next {
		cnt++
	}
	return cnt
}

// O(1)
func (l *LinkedList[T]) Count2() int {
	return l.count
}

func (l *LinkedList[T]) GetAt(idx int) *Node[T] {
	if idx >= l.Count2() {
		return nil
	}

	i := 0

	for node := l.root; node != nil; node = node.next {
		if i == idx {
			return node
		}
		i++
	}
	return nil
}

func (l *LinkedList[T]) InsertAfter(node *Node[T], value T) {
	if !l.isincluded(node) {
		return
	}

	newNode := &Node[T]{
		Value: value,
	}

	orgiNext := node.next
	node.next = newNode
	newNode.next = orgiNext

	if node == l.tail {
		l.tail = newNode
	}

	l.count++
}

func (l *LinkedList[T]) InsertBefore(node *Node[T], value T) {
	if node == l.root {
		l.PushFront(value)
		return
	}
	prevNode := l.FindPrevNode(node)
	if prevNode == nil {
		return
	}

	newNode := &Node[T]{
		Value: value,
	}

	prevNode.next = newNode
	newNode.next = node

	l.count++
}

func (l *LinkedList[T]) FindPrevNode(node *Node[T]) *Node[T] {
	inner := l.root
	for ; inner != nil; inner = inner.next {
		if inner.next == node {
			return inner
		}
	}
	return nil
}

func (l *LinkedList[T]) isincluded(node *Node[T]) bool {
	inner := l.root

	for ; inner != nil; inner = inner.next {
		if inner == node {
			return true
		}
	}
	return false
}

func (l *LinkedList[T]) PopFront() {
	if l.root == nil {
		return
	}
	l.root.next, l.root = nil, l.root.next
	if l.root == nil {
		l.tail = nil
	}
	l.count--
}

func (l *LinkedList[T]) RemoveNode(node *Node[T]) {
	if node == l.root {
		l.PopFront()
		return
	}
	prev := l.FindPrevNode(node)
	if prev == nil {
		return
	}
	if node == l.tail {
		prev.next = nil
		l.tail = prev
		l.count--
		return
	}
	l.count--
	prev.next = node.next
	node.next = nil
}
