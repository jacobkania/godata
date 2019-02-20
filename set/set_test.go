package set

import "testing"

func TestNew(t *testing.T) {
	var s Set = Set.New()

	if s.len != 0 {
		t.Errorf("Length attribute of initial Set not set to 0")
	}
	if len(s.vals) != 0 {
		t.Errorf("Initial set contains some value(s)")
	}
}
