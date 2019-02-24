package set

import "testing"

func TestNew(t *testing.T) {
	var s Set
	s.New()

	if len(s.vals) != 0 {
		t.Errorf("Initial set contains some value(s)")
	}
}

func TestToSet(t *testing.T) {

}
