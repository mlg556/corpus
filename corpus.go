/*
	Package corpus implements a 2D physical object "Corpus"
	and a 2D vector "Vector" with related mathematical operations.
*/
package corpus

import (
	"math"
)

// Corpus is a 2D physical disk with vectors for position, velocity
// acceleration, scalars for mass and radius.
type Corpus struct {
	Pos, Vel, Acc Vector
	Mass, Radius  float64
}

// Vector is a 2D vector with x and y components of type float64.
type Vector struct {
	X, Y float64
}

// MakeCorpus initialises and returns a Corpus with
// given Pos, Vel and Rad all in float64 forms.
// Mass is Rad squared, Acc is 0, 0.
func MakeCorpus(posX, posY, velX, velY, rad float64) Corpus {
	c := Corpus{}
	c.Pos = Vector{X: posX, Y: posY}
	c.Vel = Vector{X: velX, Y: velY}
	c.Radius = rad
	c.Mass = rad * rad
	c.Acc = Vector{0, 0}

	return c
}

// ApplyForce subjects the corpus to the given force by mutating its acceleration.
func (c *Corpus) ApplyForce(f Vector) {
	c.Acc.AddP(f.Div(c.Mass)) // a = F/m
}

// Update updates the given corpus by mutating its physical attributes as unit time passes.
func (c *Corpus) Update() {
	c.Vel.AddP(c.Acc) // a = dv/dt
	c.Pos.AddP(c.Vel) // v = dx/dt
	c.Acc.MultP(0)    // resets acceleration
}

// Add adds two vectors, returning a new vector.
func (a Vector) Add(b Vector) Vector {
	return Vector{a.X + b.X, a.Y + b.Y}
}

// AddI adds two vectors in place.
func (a *Vector) AddP(b Vector) {
	a.X = a.X + b.X
	a.Y = a.Y + b.Y
}

// AngleBetween returns the angle between the two vectors in radians.
func (a Vector) AngleBetween(b Vector) float64 {
	return math.Acos(a.Dot(b) / (a.Mag() * b.Mag()))
}

// Div divides the vector with a scalar(float64), returning a new vector.
func (a Vector) Div(b float64) Vector {
	return Vector{a.X / b, a.Y / b}
}

// DivP divides the vector with a scalar(float64) in place.
func (a *Vector) DivP(b float64) {
	a.X = a.X / b
	a.Y = a.Y / b
}

// Dot returns the dot product of two vectors.
func (a Vector) Dot(b Vector) float64 {
	return a.X*b.X + a.Y*b.Y
}

// Mag returns the magnitude of the vector
func (a Vector) Mag() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y)
}

// Mult multiplies the vector with a scalar(float64), returning a new vector.
func (a Vector) Mult(b float64) Vector {
	return Vector{a.X * b, a.Y * b}
}

// MultP multiplies the vector with a scalar(float64) in place.
func (a *Vector) MultP(b float64) {
	a.X = a.X * b
	a.Y = a.Y * b
}

// Norm normalizes a vector, returning a new vector.
func (a Vector) Norm() Vector {
	return a.Div(a.Mag())
}

// NormP normalizes a vector in place.
func (a *Vector) NormP() {
	a.DivP(a.Mag())
}

// Sub substracts a from b, returning a new vector.
func (a Vector) Sub(b Vector) Vector {
	return Vector{a.X - b.X, a.Y - b.Y}
}

// SubP substracts a from b in place.
func (a *Vector) SubP(b Vector) {
	a.X = a.X - b.X
	a.Y = a.Y - b.Y
}
