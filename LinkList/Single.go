package linklist

type SNode[T any] struct {
	next  *SNode[T]
	value T
}
type LinkedList[T any] struct {
	root *SNode[T]
	tail *SNode[T]
}

// add node back
func (l *LinkedList[T]) PushBack(value T) {
	node := &SNode[T]{
		value: value,
	}
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

func (l LinkedList[T]) back() *SNode[T] {
	return l.tail
}
