package godata

import "sync"

type node struct {
	value interface{}
	prev  *node
}

// Stack data structure
type Stack struct {
	topNode *node
	lock    sync.Mutex
}

// New creates a new stack
func (s *Stack) New() {
	s = &Stack{nil, sync.Mutex{}}
}

// Push adds a value to the top of the stack
func (s *Stack) Push(value interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	newNode := node{value, s.topNode}
	s.topNode = &newNode
}

// Pop removes the value from the top of the stack
func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	top := *s.topNode
	s.topNode = s.topNode.prev
	return top.value
}
