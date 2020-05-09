package myq

import (
	"container/list"
)

// Queue is a queue
type Queue interface {
	Front() *list.Element
	Len() int
	Add(interface{})
	Remove() interface{}
}

type queueImpl struct {
	*list.List
}

func (q *queueImpl) Len() int {
	return q.List.Len()
}

func (q *queueImpl) Add(v interface{}) {
	q.PushBack(v)
}

func (q *queueImpl) Remove() interface{} {
	e := q.Front()
	q.List.Remove(e)
	return e
}

// New is a new instance of a Queue
func NewQ() Queue {
	return &queueImpl{list.New()}
}
