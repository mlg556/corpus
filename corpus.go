/*
	Package corpus implements a 2D physical object "Corpus"
	and a 2D vector "Vector" with related mathematical operations.
 */
package corpus

import (
	"math"
)

// Corpus is a 2D physical disk with vectors for position, velocity
// acceleration, scalars for mass and radius, and a boolean state for immobility
type Corpus struct {
	Pos, Vel, Acc Vector
	Mass, Radius float64
	Immob bool
}

// Vector is a 2D vector with x and y components of type float64.
type Vector struct {
	X, Y float64
}

// Add adds two vectors, returning a new vector.
func (a Vector) Add(b Vector) Vector {
	return Vector{a.X + b.X, a.Y + b.Y}
}

// AddI adds two vectors in place.
func (a *Vector) AddI(b Vector) {
	a.X = a.X + b.X
	a.Y = a.Y + b.Y
}

// AngleBetween returns the angle between the two vectors in radians.
func (a Vector) AngleBetween(b Vector) float64 {
	return math.Acos(a.Dot(b) / (a.Mag()*b.Mag()))
}

// Div divides the vector with a scalar(float64), returning a new vector.
func (a Vector) Div(b float64) Vector {
	return Vector{a.X / b, a.Y / b}
}

// DivI divides the vector with a scalar(float64) in place.
func (a *Vector) DivI(b float64) {
	a.X = a.X / b
	a.Y = a.Y / b
}

// Dot returns the dot product of two vectors.
func (a Vector) Dot(b Vector) float64 {
	return a.X * b.X + a.Y * b.Y
}

// Mag returns the magnitude of the vector
func (a Vector) Mag() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y)
}

// Mult multiplies the vector with a scalar(float64), returning a new vector.
func (a Vector) Mult(b float64) Vector {
	return Vector{a.X * b, a.Y * b}
}

// MultI multiplies the vector with a scalar(float64) in place.
func (a *Vector) MultI(b float64) {
	a.X = a.X * b
	a.Y = a.Y * b
}

// Norm normalizes a vector, returning a new vector.
func (a Vector) Norm() Vector {
	return a.Div(a.Mag())
}

// NormI normalizes a vector in place.
func (a *Vector) NormI() {
	a.DivI(a.Mag())
}

// Sub substracts a from b, returning a new vector.
func (a Vector) Sub(b Vector) Vector {
	return Vector{a.X - b.X, a.Y - b.Y}
}

// SubI substracts a from b in place.
func (a *Vector) SubI(b Vector) {
	a.X = a.X - b.X
	a.Y = a.Y - b.Y
}













