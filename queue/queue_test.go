package queue

import "testing"

func TestNew(t *testing.T) {
	var q Queue
	q.New()

	if len(q.data) != 0 {
		t.Errorf("Initial data is not empty, got %v", q.data)
	}
	if q.len != 0 {
		t.Errorf("Initial length value was not 0, got %v", q.len)
	}
}

func TestPush(t *testing.T) {
	var q Queue
	q.New()

	q.Push(1)

	if q.Peek() != 1 {
		t.Errorf("Failed to add node to stack, got %v", q.Peek())
	}
}

func TestPushSeveral(t *testing.T) {
	var q Queue
	q.New()

	q.Push(1)
	q.Push(2)
	q.Push(3)

	if q.Peek() != 1 {
		t.Errorf("Front value altered, got %v", q.Peek())
	}
}

func TestPull(t *testing.T) {
	var q Queue
	q.New()

	q.Push(1)
	q.Push(2)
	q.Push(3)

	returnedVal := q.Pull()

	if returnedVal != 1 {
		t.Errorf("Pulled value not returned successfully, got %v", returnedVal)
	}
}

func TestPullEmptyQueue(t *testing.T) {
	var q Queue
	q.New()

	returnedVal := q.Pull()

	if returnedVal != nil {
		t.Errorf("Somehow didn't return nothing, got %v", returnedVal)
	}
}

func TestPeek(t *testing.T) {
	var q Queue
	q.New()

	q.Push(1)
	q.Push(2)
	q.Push(3)

	returnedVal := q.Peek()

	if returnedVal != 1 {
		t.Errorf("Peeked value not returned successfully, got %v", returnedVal)
	}

	if len(q.data) != 3 {
		t.Errorf("Data was modified during peek")
	}
}

func TestPeekEmptyQueue(t *testing.T) {
	var q Queue
	q.New()

	returnedVal := q.Peek()

	if returnedVal != nil {
		t.Errorf("Somehow didn't return nothing, got %v", returnedVal)
	}
}

func TestLen_Initial(t *testing.T) {
	var q Queue
	q.New()

	if q.Len() != 0 {
		t.Errorf("Initial length incorrect, got %v", q.Len())
	}
}

func TestLen_IncreasesWithPush(t *testing.T) {
	var q Queue
	q.New()

	q.Push(1)
	q.Push(2)

	if q.Len() != 2 {
		t.Errorf("Length incorrect, got %v", q.Len())
	}
}

func TestLen_DecreasesWithPop(t *testing.T) {
	var q Queue
	q.New()

	q.Push(1)
	q.Push(2)

	q.Pull()

	if q.Len() != 1 {
		t.Errorf("Length incorrect, got %v", q.Len())
	}
}
