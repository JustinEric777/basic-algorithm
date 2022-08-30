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
		next  *node
	}
)

func NewQueue() *Queue {
	return &Queue{nil, nil, 0, &sync.RWMutex{}}
}

func (this *Queue) Push(value interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()

	node := &node{value, this.top, nil}
	if this.top == nil {
		this.top = node
		this.tail = this.top
	} else {
		node.prev = this.tail
		this.tail.next = node
		this.tail = node
	}
	this.length++
}

func (this *Queue) Pop() interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()

	if this.length == 0 {
		return nil
	}
	curNode := this.top
	if this.top.next == nil {
		this.top = nil
	} else {
		this.top = this.top.next
		this.top.prev.next = nil
		this.top.prev = nil
	}

	this.length--

	return curNode.value
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
