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
	assert.Equal(t, 1, l.Front().value)
	assert.Equal(t, 1, l.Back().value)

	l.PushBack(2)
	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().value)
	assert.Equal(t, 2, l.Back().value)
}

func TestPushFront(t *testing.T) {
	var l LinkedList[int]

	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)

	l.PushFront(1)
	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().value)
	assert.Equal(t, 1, l.Back().value)

	l.PushFront(2)
	assert.NotNil(t, l.root)
	assert.Equal(t, 2, l.Front().value)
	assert.Equal(t, 1, l.Back().value)

	l.PushFront(4) // node) 4->2->1
	assert.NotNil(t, l.root)
	assert.Equal(t, 4, l.Front().value)
	assert.Equal(t, 1, l.Back().value)

	assert.Equal(t, 3, l.CountNode())
	assert.Equal(t, 3, l.CountNode2())

	assert.Equal(t, 4, l.GetAt(0).value)
	assert.Equal(t, 2, l.GetAt(1).value)
	assert.Equal(t, 1, l.GetAt(2).value)

}

func TestInsertAfter(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	node := l.GetAt(1)
	l.InsertAfter(node, 4)

	assert.Equal(t, 4, l.CountNode2())
	assert.Equal(t, 4, l.GetAt(2).value)
	assert.Equal(t, 3, l.Back().value)

	tempNode := &SNode[int]{
		value: 100,
	}
	l.InsertAfter(tempNode, 200)

	assert.Equal(t, 4, l.CountNode())
	assert.Equal(t, 4, l.CountNode2())
}
