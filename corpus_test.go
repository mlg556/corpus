package corpus

import (
	"testing"
)

func TestVector_Add(t *testing.T) {
	a := Vector{1, 1}
	b := Vector{2, 2}

	res := Vector{3, 3}

	if a.Add(b) != res {
		t.Error("WRONG!!")
	}

}

func TestVector_AddI(t *testing.T) {
	a := Vector{1, 1}
	b := Vector{2, 2}
	a.AddI(b)
	res := Vector{3, 3}

	if a != res {
		t.Error("WRONG!!")
	}

}

func TestVector_Mag(t *testing.T) {
	a := Vector{3, 4}
	res := 5.0

	if a.Mag() != res {
		t.Error("WRONG!!")
	}
}