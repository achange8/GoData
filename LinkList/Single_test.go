package linklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushBack(t *testing.T) {
	var l LinkedList[int]

	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)
	l.PushBack(1)

	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	l.PushBack(2)
	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 2, l.Back().Value)
}

func TestPushFront(t *testing.T) {
	var l LinkedList[int]

	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)

	l.PushFront(1)
	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	l.PushFront(2)
	assert.NotNil(t, l.root)
	assert.Equal(t, 2, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	l.PushFront(4) // node) 4->2->1
	assert.NotNil(t, l.root)
	assert.Equal(t, 4, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	assert.Equal(t, 3, l.Count())
	assert.Equal(t, 3, l.Count2())

	assert.Equal(t, 4, l.GetAt(0).Value)
	assert.Equal(t, 2, l.GetAt(1).Value)
	assert.Equal(t, 1, l.GetAt(2).Value)

}

func TestInsertAfter(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	node := l.GetAt(1)
	l.InsertAfter(node, 4)

	assert.Equal(t, 4, l.Count2())
	assert.Equal(t, 4, l.GetAt(2).Value)
	assert.Equal(t, 3, l.Back().Value)

	tempNode := &Node[int]{
		Value: 100,
	}
	l.InsertAfter(tempNode, 200)

	assert.Equal(t, 4, l.Count())
	assert.Equal(t, 4, l.Count2())
}

func TestInsertBefore(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3) //1423
	assert.Equal(t, 3, l.Count())
	assert.Equal(t, 3, l.Count2())

	node := l.GetAt(1)
	l.InsertBefore(node, 4)

	assert.Equal(t, 4, l.Count())
	assert.Equal(t, 4, l.Count2())
	assert.Equal(t, 1, l.GetAt(0).Value)
	assert.Equal(t, 4, l.GetAt(1).Value)
	assert.Equal(t, 2, l.GetAt(2).Value)
	assert.Equal(t, 3, l.GetAt(3).Value)
	assert.Equal(t, 3, l.Back().Value)

	tempNode := &Node[int]{
		Value: 100,
	}
	l.InsertBefore(tempNode, 200)

	assert.Equal(t, 4, l.Count())
	assert.Equal(t, 4, l.Count2())
}

func TestPopFront(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3) //123
	l.PopFront()  //23
	assert.Equal(t, 2, l.Count())
	assert.Equal(t, 2, l.Count2())
	assert.Equal(t, 2, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)
	l.PopFront() //3
	assert.Equal(t, 1, l.Count())
	assert.Equal(t, 1, l.Count2())
	assert.Equal(t, 3, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)

	l.PopFront() //nil
	assert.Equal(t, 0, l.Count())
	assert.Equal(t, 0, l.Count2())
	assert.Nil(t, l.Front())
	assert.Nil(t, l.Back())

}

func TestRemoveNode(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)            //1234
	l.RemoveNode(l.GetAt(1)) //134
	assert.Equal(t, 3, l.Count())
	assert.Equal(t, 3, l.Count2())
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 4, l.Back().Value)

	l.RemoveNode(l.GetAt(2)) //13
	assert.Equal(t, 2, l.Count())
	assert.Equal(t, 2, l.Count2())
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)

	l.RemoveNode(l.GetAt(1)) //1
	assert.Equal(t, 1, l.Count())
	assert.Equal(t, 1, l.Count2())
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	l.RemoveNode(&Node[int]{
		Value: 100,
	})
	assert.Equal(t, 1, l.Count())
	assert.Equal(t, 1, l.Count2())
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	l.RemoveNode(l.GetAt(0)) //nil
	assert.Equal(t, 0, l.Count())
	assert.Equal(t, 0, l.Count2())
	assert.Nil(t, l.Front())
	assert.Nil(t, l.Back())

}
