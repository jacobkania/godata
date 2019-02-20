package set

// Set set data structure
type Set struct {
	vals map[interface{}]bool
	len  int
}

// New creates a new Set
func (s Set) New() Set {
	return Set{make(map[interface{}]bool), 0}
}

// Add adds a value to a set
func (s *Set) Add(val interface{}) {
	s.vals[val] = true
}

// AddMany adds many values to a set
func (s *Set) AddMany(vals []interface{}) {
	for val := range vals {
		s.Add(val)
	}
}

// Contains returns true if the given value is in the set
func (s *Set) Contains(val interface{}) bool {
	if _, ok := s.vals[val]; ok {
		return true
	}
	return false
}

// Remove removes a value from the set
// returns false if val doesn't exist in the set
func (s *Set) Remove(val interface{}) bool {
	if _, ok := s.vals[val]; !ok {
		return false
	}

	delete(s.vals, val)

	return true
}

// Len returns number of items in the set
func (s *Set) Len() int {
	return s.len
}
