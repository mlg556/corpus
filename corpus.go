/*
	Package corpus implements a 2D physical object "Corpus"
	and a 2D vector "Vector" with related mathematical operations.
*/
package corpus

import (
	"math"
)

// Corpus is a 2D physical disk with vectors for position, velocity
// acceleration, scalars for mass and radius and a boolean Immaterial
// for things simply to be drawn.
type Corpus struct {
	Pos, Vel, Acc        Vector
	Mass, Charge, Radius float64
	Immaterial           bool
}

// Vector is a 2D vector with x and y components of type float64.
type Vector struct {
	X, Y float64
}

// MakeCorpus initialises and returns a Corpus with
// given Pos, Vel and Rad all in float64 forms.
// By default, Acc is 0,0 and Immaterial is false.
func MakeCorpus(posX, posY, velX, velY, mass, charge, rad float64) Corpus {
	c := Corpus{}
	c.Pos = Vector{X: posX, Y: posY}
	c.Vel = Vector{X: velX, Y: velY}
	c.Mass = mass
	c.Charge = charge
	c.Radius = rad
	c.Acc = Vector{0, 0}
	c.Immaterial = false

	return c
}

// IsInter checks if two corpi intersect each other.
func (c Corpus) IsInter(cp *Corpus) bool {
	return c.Pos.Dist(cp.Pos) <= c.Radius+cp.Radius
}

// ApplyForce subjects the corpus to the given force by mutating its acceleration.
// By Newton's 2nd law "F = m*a".
func (c *Corpus) ApplyForce(f Vector) {
	c.Acc.AddP(f.Div(c.Mass)) // a = F/m
}

// Bounce bounces the corpus off windows boundaries given by width and height.
// Also prevents intersections by directly mutating position.
func (c *Corpus) Bounce(width, height float64) {
	posX := c.Pos.X
	posY := c.Pos.Y
	rad := c.Radius

	if posX > width-rad {
		c.Pos.X = width - rad
		c.Vel.X = -c.Vel.X
	}
	if posX < rad {
		c.Pos.X = rad
		c.Vel.X = -c.Vel.X
	}
	if posY > height-rad {
		c.Pos.Y = height - rad
		c.Vel.Y = -c.Vel.Y
	}
	if posY < rad {
		c.Pos.Y = rad
		c.Vel.Y = -c.Vel.Y
	}
}

// Collide collides the given corpus with the given corpi in the slice.
// Mutates velocities to preserve momentum and kinetic energy.
// Also prevents intersections by directly mutating positions.
// Solving the equation for a 2D elastic collision yields:
// v1' = v1 - (2*m2/(m1+m2)) * (<v1-v2, x1-x2>)/(||x1-x2||^2) * (x1-x2)
// v2' = v2 - (2*m1/(m1+m2)) * (<v2-v1, x2-x1>)/(||x2-x1||^2) * (x2-x1)
func (c *Corpus) Collide(corpi []Corpus) {
	for idx := range corpi {
		cp := &corpi[idx]
		if !c.Immaterial && !cp.Immaterial {
			dist := c.Pos.Dist(cp.Pos)
			// There is a collision.
			if c.IsInter(cp) {
				//Intersection
				displace := c.Pos.Sub(cp.Pos).SetMag((c.Radius + cp.Radius - dist) / 2)
				c.Pos.AddP(displace)
				cp.Pos.AddP(displace.Mult(-1))
				// Momentum
				cVelP := c.Vel.Sub(c.Pos.Sub(cp.Pos).Mult((2 * cp.Mass / (c.Mass + cp.Mass)) * c.Vel.Sub(cp.Vel).Dot(c.Pos.Sub(cp.Pos)) / c.Pos.DistSq(cp.Pos)))
				cp.Vel = cp.Vel.Sub(cp.Pos.Sub(c.Pos).Mult((2 * c.Mass / (c.Mass + cp.Mass)) * cp.Vel.Sub(c.Vel).Dot(cp.Pos.Sub(c.Pos)) / c.Pos.DistSq(cp.Pos)))
				c.Vel = cVelP
			}
		}
	}
}

// Gravitate calculates and applies the gravitational force
// between the given corpus and all of the rest corpi.
// Using Newton's law of universal gravitation:
// F = G*m1*m2/r^2
func (c *Corpus) Gravitate(corpi []Corpus, G float64) {
	for idx := range corpi {
		cp := &corpi[idx]
		if !c.Immaterial && !cp.Immaterial {
			dist := c.Pos.Dist(cp.Pos)
			if dist+2 >= c.Radius+cp.Radius {
				force := cp.Pos.Sub(c.Pos).Mult(G * c.Mass * cp.Mass / (dist * dist)).Div(dist)
				c.ApplyForce(force)
				cp.ApplyForce(force.Mult(-1))
			}
		}
	}
}

// Coulomb calculates and applies the electrostatic force
// between the given corpus and all of the rest corpi.
// Using Coulomb's law:
// F = k*q1*q2/r^2
func (c *Corpus) Coulomb(corpi []Corpus, k float64) {
	for idx := range corpi {
		cp := &corpi[idx]
		if !c.Immaterial && !cp.Immaterial {
			dist := c.Pos.Dist(cp.Pos)
			if dist+2 >= c.Radius+cp.Radius {
				force := cp.Pos.Sub(c.Pos).Mult(1 * c.Charge * cp.Charge / (dist * dist)).Div(dist)
				c.ApplyForce(force.Mult(-1))
				cp.ApplyForce(force)
			}
		}
	}
}

// Update updates the given corpus by mutating its physical attributes as unit time passes.
func (c *Corpus) Update() {
	if !c.Immaterial {
		c.Vel.AddP(c.Acc) // a = dv/dt
		c.Pos.AddP(c.Vel) // v = dx/dt
		c.Acc.MultP(0)    // resets acceleration
	}
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

// Dist returns the distance between the two vectors as a float64 number.
func (a Vector) Dist(b Vector) float64 {
	return a.Sub(b).Mag()
}

// DistSq returns the distance between the two vectors, squared as a float64 number.
func (a Vector) DistSq(b Vector) float64 {
	return a.Sub(b).MagSq()
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

// MagSq returns the magnitude of the vector, squared.
func (a Vector) MagSq() float64 {
	return a.X*a.X + a.Y*a.Y
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

// SetMag returns a new vector in the same direction with given magnitude.
func (a Vector) SetMag(mag float64) Vector {
	return a.Norm().Mult(mag)
}

// SetMagP mutates the magnitude of the vector, keeping the direction.
func (a *Vector) SetMagP(mag float64) {
	a.NormP()
	a.MultP(mag)
}

// Sub subtracts a from b, returning a new vector.
func (a Vector) Sub(b Vector) Vector {
	return Vector{a.X - b.X, a.Y - b.Y}
}

// SubP subtracts a from b in place.
func (a *Vector) SubP(b Vector) {
	a.X = a.X - b.X
	a.Y = a.Y - b.Y
}
