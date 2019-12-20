package vector

import (
	"math"
)

// Vector2f represents two dimensional vectors
type Vector2f struct {
	X float64
	Y float64
}

// NewVector2f creates a new two dimensional vector
func NewVector2f(x, y float64) *Vector2f {
	return &Vector2f{
		X: x,
		Y: y,
	}
}

var (
	// UpUnit2f points in the positive direction on the y-axis
	UpUnit2f = NewVector2f(0, 1)
	// RightUnit2f points in the positive direciton on the x-axis
	RightUnit2f = NewVector2f(1.0, 0.0)
	// DownUnit2f points in the negative direction on the y-axis
	DownUnit2f = NewVector2f(0.0, -1.0)
	// LeftUnit2f points in the negative direciton on the x-axis
	LeftUnit2f = NewVector2f(-1.0, 0.0)
)

// Units configures the up, down, left and right unit vectors
func Units(up, right *Vector2f) {
	UpUnit2f = up
	RightUnit2f = right
	DownUnit2f = Multiply2f(UpUnit2f, NewVector2f(0, -1))
	LeftUnit2f = Multiply2f(RightUnit2f, NewVector2f(-1, 0))
}

// Zero2f zero Vector
func Zero2f() *Vector2f {
	return NewVector2f(0, 0)
}

// Equals compares two Vector2f
func (a *Vector2f) Equals(b *Vector2f) bool {
	return a.X == b.X && a.Y == b.Y
}

// Set updates the values
func (a *Vector2f) Set(x, y float64) {
	a.X = x
	a.Y = y
}

// SetX updates the x value
func (a *Vector2f) SetX(x float64) {
	a.X = x
}

// SetY updates the y value
func (a *Vector2f) SetY(y float64) {
	a.Y = y
}

// Add2f two Vector2f and create a new one with the result
func Add2f(a, b *Vector2f) *Vector2f {
	return NewVector2f(a.X+b.X, a.Y+b.Y)
}

// Add a Vector2f to this Vector2f
func (a *Vector2f) Add(b *Vector2f) {
	a.X += b.X
	a.Y += b.Y
}

// Subtract2f a Vector2f from another and create a new one with the result
func Subtract2f(a, b *Vector2f) *Vector2f {
	return NewVector2f(a.X-b.X, a.Y-b.Y)
}

// Subtract a Vector2f from this Vector2f
func (a *Vector2f) Subtract(b *Vector2f) {
	a.X -= b.X
	a.Y -= b.Y
}

// Multiply2f a Vector2f with another and create a new one with the result
func Multiply2f(a, b *Vector2f) *Vector2f {
	return NewVector2f(a.X*b.X, a.Y*b.Y)
}

// Multiply this Vector2f by another
func (a *Vector2f) Multiply(b *Vector2f) {
	a.X *= b.X
	a.Y *= b.Y
}

// Divide2f a Vector2f with another and create a new one with the result
func Divide2f(a, b *Vector2f) *Vector2f {
	return NewVector2f(a.X/b.X, a.Y/b.Y)
}

// Divide this Vector2f by another
func (a *Vector2f) Divide(b *Vector2f) {
	a.X /= b.X
	a.Y /= b.Y
}

// Scale2f a Vector2f by a constant scale factor and create a new one with the result
func Scale2f(a *Vector2f, scale float64) *Vector2f {
	return NewVector2f(a.X*scale, a.Y*scale)
}

// Scale this Vector2f by a constant
func (a *Vector2f) Scale(scale float64) {
	a.X *= scale
	a.Y *= scale
}

// Magnitude of the Vector2f
func (a *Vector2f) Magnitude() float64 {
	return math.Sqrt((a.X * a.X) + (a.Y * a.Y))
}

// Normalize2f and return a new Vector2f
func Normalize2f(a *Vector2f) *Vector2f {
	return Scale2f(a, 1.0/a.Magnitude())
}

// Normalize the vector (convert to unit vector)
func (a *Vector2f) Normalize() {
	a.Scale(1.0 / a.Magnitude())
}

// DotProduct2f of two Vector2f, returned as a new Vector2f
func DotProduct2f(a, b *Vector2f) float64 {
	return (a.X * b.X) + (a.Y * b.Y)
}

// Determinant of two Vector2f, returned as a new Vector2f
func Determinant(a, b *Vector2f) float64 {
	return (a.X * b.Y) + (a.Y * b.X)
}

// AngleBetween two Vector2f
func AngleBetween(a, b *Vector2f) float64 {
	dotProduct := DotProduct2f(a, b)
	determinant := Determinant(a, b)
	return math.Atan2(determinant, dotProduct)
}
