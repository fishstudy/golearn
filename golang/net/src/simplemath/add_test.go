package simplemath

import "testing"

func TestAdd1(t *testing.T) {
	r := Add(1, 2)
	if r != 3 {
		t.Error("Add(1,2) fail")
	}
}
