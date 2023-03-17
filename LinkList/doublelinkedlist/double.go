package doublelinklist

type Node[T any] struct {
	next  *Node[T]
	prev  *Node[T]
	Value T
}

type LinkedList[T any] struct {
	root  *Node[T]
	tail  *Node[T]
	count int
}

func (l *LinkedList[T]) PushBack(val T) {
	n := &Node[T]{
		Value: val,
	}
	if l.tail == nil {
		l.root = n
		l.tail = n
		l.count++
		return
	}
	l.tail.next = n
	n.prev = l.tail
	l.tail = n
	l.count++

}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.root
}

func (l *LinkedList[T]) Back() *Node[T] {
	return l.tail
}

func (l *LinkedList[T]) Count() int {
	return l.count
}

func (l *LinkedList[T]) GetAt() *Node[T] {

}
