package corpus

import (
	"math"
	"testing"
)

func TestMakeCorpus(t *testing.T) {
	posX := 1.0
	posY := 2.0
	velX := 3.0
	velY := 4.0
	rad := 5.0
	c := MakeCorpus(posX, posY, velX, velY, rad)

	if c.Pos.X != posX {
		t.Error("WRONG POSX")
	}

	if c.Pos.Y != posY {
		t.Error("WRONG POSY")
	}

	if c.Vel.X != velX {
		t.Error("WRONG VELX")
	}

	if c.Vel.Y != velY {
		t.Error("WRONG VELY")
	}

	if c.Radius != rad {
		t.Error("WRONG RAD")
	}

}

func TestCorpus_ApplyForce(t *testing.T) {
	pos := Vector{0, 0}
	vel := Vector{0, 0}
	acc := Vector{0, 0}

	mass := 1.0
	radius := 1.0

	c := Corpus{pos, vel, acc, mass, radius}

	c.ApplyForce(Vector{1, 1})

	res := Vector{1, 1}

	if c.Acc != res {
		t.Error("WRONG !!")
	}
}

func TestCorpus_Update(t *testing.T) {
	pos := Vector{0, 0}
	vel := Vector{1, 1}
	acc := Vector{1, 1}

	mass := 1.0
	radius := 1.0

	c := Corpus{pos, vel, acc, mass, radius}
	c.Update()

	resPos := Vector{2, 2}
	resVel := Vector{2, 2}
	resAcc := Vector{0, 0}

	if c.Pos != resPos {
		t.Error("WRONG POS!!")
	}

	if c.Vel != resVel {
		t.Error("WRONG VEL")
	}

	if c.Acc != resAcc {
		t.Error("WRONG ACC")
	}
}

func TestVector_Add(t *testing.T) {
	a := Vector{1, 1}
	b := Vector{2, 2}

	res := Vector{3, 3}

	if a.Add(b) != res {
		t.Error("WRONG!!")
	}

}

func TestVector_AddP(t *testing.T) {
	a := Vector{1, 1}
	b := Vector{2, 2}
	a.AddP(b)
	res := Vector{3, 3}

	if a != res {
		t.Error("WRONG!!")
	}

}

func TestVector_AngleBetween(t *testing.T) {
	a := Vector{1, 0}
	b := Vector{0, 1}

	res := 2 * math.Pi

	if a.AngleBetween(b)-res > 0.001 {
		t.Error("WRONG!!")
	}
}

func TestVector_Div(t *testing.T) {
	a := Vector{2, 2}
	res := Vector{1, 1}

	if a.Div(2) != res {
		t.Error("WRONG !!")
	}
}

func TestVector_DivP(t *testing.T) {
	a := Vector{2, 2}
	res := Vector{1, 1}

	a.DivP(2)

	if a != res {
		t.Error("WRONG !!")
	}
}

func TestVector_Dot(t *testing.T) {
	a := Vector{1, 2}
	b := Vector{3, 4}
	res := 11.0

	if a.Dot(b) != res {
		t.Error("WRONG")
	}
}

func TestVector_Mag(t *testing.T) {
	a := Vector{3, 4}
	res := 5.0

	if a.Mag() != res {
		t.Error("WRONG!!")
	}
}

func TestVector_Mult(t *testing.T) {
	a := Vector{3, 4}
	b := 3.0

	res := Vector{9, 12}

	if a.Mult(b) != res {
		t.Error("WRONG !!")
	}
}

func TestVector_MultP(t *testing.T) {
	a := Vector{3, 4}
	b := 3.0
	a.MultP(b)

	res := Vector{9, 12}

	if a != res {
		t.Error("WRONG !!")
	}
}

func TestVector_Norm(t *testing.T) {
	a := Vector{10, 0}
	res := Vector{1, 0}

	if a.Norm() != res {
		t.Error("WRONG !!")
	}
}

func TestVector_NormP(t *testing.T) {
	a := Vector{10, 0}
	res := Vector{1, 0}

	a.NormP()

	if a != res {
		t.Error("WRONG !!")
	}
}

func TestVector_Sub(t *testing.T) {
	a := Vector{3, 4}
	b := Vector{1, 2}

	res := Vector{2, 2}

	if a.Sub(b) != res {
		t.Error("WRONG !!")
	}
}

func TestVector_SubP(t *testing.T) {
	a := Vector{3, 4}
	b := Vector{1, 2}

	res := Vector{2, 2}

	a.SubP(b)

	if a != res {
		t.Error("WRONG !!")
	}
}
