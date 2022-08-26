package queue

import "sync"

type (
	Queue struct {
		top    *node
		tail   *node
		length int
		lock   *sync.RWMutex
	}
	node struct {
		value interface{}
		prev  *node
	}
)

func NewQueue() *Queue {
	return &Queue{nil, nil, 0, &sync.RWMutex{}}
}

func (this *Queue) Push(value interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()

	node := &node{value, this.tail.prev}
	this.tail = node
	this.length++
}

func (this *Queue) Pop() interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()

	if this.length == 0 {
		return nil
	}
	value := this.top.value
	this.top = this.top.prev
	this.length--

	return value
}

func (this *Queue) Len() int {
	return this.length
}

func (this *Queue) Peek() interface{} {
	if this.length == 0 {
		return nil
	}

	return this.top.value
}

func (this *Queue) Tail() interface{} {
	if this.length == 0 {
		return nil
	}

	return this.tail.value
}
