package stack

import (
	"sync"
)

type (
	Stack struct {
		top    *node
		length int
		lock   *sync.RWMutex
	}
	node struct {
		value interface{}
		prev  *node
	}
)

func NewStack() *Stack {
	return &Stack{nil, 0, &sync.RWMutex{}}
}

func (this *Stack) Push(val interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()

	node := &node{prev: this.top, value: val}
	this.top = node
	this.length++
}

func (this *Stack) Pop() interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()

	if this.length == 0 {
		return nil
	}

	value := this.top.value
	this.length--
	this.top = this.top.prev

	return value
}

func (this *Stack) Len() int {
	return this.length
}

func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}

	return this.top.value
}
