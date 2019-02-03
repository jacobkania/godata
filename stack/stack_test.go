package godata

import (
	"testing"
)

func TestNew(t *testing.T) {
	var s Stack
	s.New()

	if s.topNode != nil {
		t.Errorf("topNode was not initialized to nil, got %v", s.topNode)
	}
}

func TestPush(t *testing.T) {
	var s Stack
	s.New()

	s.Push(1)

	if s.topNode.value != 1 {
		t.Errorf("Failed to add node to stack, got %v", s.topNode.value)
	}
}

func TestPushSeveral(t *testing.T) {
	var s Stack
	s.New()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.topNode.value != 3 {
		t.Errorf("Third value added not on top of stack, got %v", s.topNode.value)
	}
}

func TestPop(t *testing.T) {
	var s Stack
	s.New()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	returnedVal := s.Pop()

	if s.topNode.value != 2 {
		t.Errorf("Third value not removed successfully, got %v", s.topNode.value)
	}

	if returnedVal != 3 {
		t.Errorf("Popped value not returned successfully, got %v", returnedVal)
	}
}

func TestPeek(t *testing.T) {
	var s Stack
	s.New()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	returnedVal := s.Peek()

	if s.topNode.value != 3 {
		t.Errorf("Top of stack was modified, got %v", s.topNode.value)
	}

	if returnedVal != 3 {
		t.Errorf("Failed to peek at top value, got %v", returnedVal)
	}
}

func TestLenInitial(t *testing.T) {
	var s Stack
	s.New()

	if s.Len() != 0 {
		t.Errorf("Initial length incorrect, got %v", s.Len())
	}
}

func TestLen(t *testing.T) {
	var s Stack
	s.New()

	s.Push(1)
	s.Push(2)

	if s.Len() != 2 {
		t.Errorf("Length incorrect, got %v", s.Len())
	}
}

func TestLenAfterPop(t *testing.T) {
	var s Stack
	s.New()

	s.Push(1)
	s.Push(2)

	s.Pop()

	if s.Len() != 1 {
		t.Errorf("Length incorrect, got %v", s.Len())
	}
}
