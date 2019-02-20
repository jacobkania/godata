package stack

import "sync"

type node struct {
	value interface{}
	prev  *node
}

// Stack data structure
type Stack struct {
	topNode *node
	lock    sync.Mutex
	len     int
}

// New creates a new stack
func (s *Stack) New() {
	s = &Stack{nil, sync.Mutex{}, 0}
}

// Push adds a value to the top of the stack
func (s *Stack) Push(value interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	newNode := node{value, s.topNode}
	s.topNode = &newNode
	s.len++
}

// Pop removes and returns the value from the top of the stack
func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.len == 0 {
		return nil
	}

	topVal := s.topNode.value
	s.topNode = s.topNode.prev
	s.len--
	return topVal
}

// Peek returns (without removing) the value from the top of the stack
func (s *Stack) Peek() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.len == 0 {
		return nil
	}

	return s.topNode.value
}

// Len gives the size of the stack
func (s *Stack) Len() int {
	return s.len
}
