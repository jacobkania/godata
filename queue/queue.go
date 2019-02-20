package queue

import "sync"

// Queue data structure
type Queue struct {
	data []interface{}
	lock sync.Mutex
	len  int
}

// New creates a new queue
func (q *Queue) New() {
	q = &Queue{[]interface{}{}, sync.Mutex{}, 0}
}

// Push adds a value to the back of the queue
func (q *Queue) Push(value interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.data = append(q.data, value)
	q.len++
}

// Pull removes and returns the value from the front of the queue
func (q *Queue) Pull() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.len == 0 {
		return nil
	}

	frontVal := q.data[0]
	q.data = q.data[1:]
	q.len--
	return frontVal
}

// Peek returns (without removing) the value from the front of the queue
func (q *Queue) Peek() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.len == 0 {
		return nil
	}

	return q.data[0]
}

// Len gives the size of the queue
func (q *Queue) Len() int {
	return q.len
}
