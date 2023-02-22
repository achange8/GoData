package linklist

type SNode[T any] struct {
	next  *SNode[T]
	value T
}
type LinkedList[T any] struct {
	root *SNode[T]
	tail *SNode[T]

	count int
}

// add node back
func (l *LinkedList[T]) PushBack(value T) {
	node := &SNode[T]{
		value: value,
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
	node := &SNode[T]{
		value: value,
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

func (l LinkedList[T]) Front() *SNode[T] {
	return l.root
}

func (l LinkedList[T]) Back() *SNode[T] {
	return l.tail
}

// O(N)
func (l LinkedList[T]) CountNode() int {
	node := l.root
	cnt := 0

	for ; node != nil; node = node.next {
		cnt++
	}
	return cnt
}

// O(1)
func (l LinkedList[T]) CountNode2() int {
	return l.count
}

func (l LinkedList[T]) GetAt(idx int) *SNode[T] {
	if idx >= l.CountNode2() {
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

func (l *LinkedList[T]) InsertAfter(node *SNode[T], value T) {
	if !l.isinclude(node) {
		return
	}

	newNode := &SNode[T]{
		value: value,
	}

	orgiNext := node.next
	node.next = newNode
	newNode.next = orgiNext

	l.count++
}

func (l LinkedList[T]) InsertBefore(node *SNode[T], value T) {
	if node == l.root {
		l.PushFront(value)
		return
	}
	PrevNode := l.findPrevNode(node)
	if PrevNode == nil {
		return
	}

	newNode := &SNode[T]{
		value: value,
	}

}

func (l LinkedList[T]) findPrevNode(node *SNode[T]) *SNode[T] {
	inner := l.root
	for ; inner != nil; inner = inner.next {
		if inner.next == node {
			return inner
		}
	}
	return nil
}

func (l LinkedList[T]) isinclude(node *SNode[T]) bool {
	inner := l.root

	for ; inner != nil; inner = inner.next {
		if inner == node {
			return true
		}
	}
	return false
}
