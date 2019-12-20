package vector

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

const EqualityEpsilon float64 = 0.000000001

func TestUnitVectors(t *testing.T) {
	Units(NewVector2f(0, 1), NewVector2f(1, 0))

	assert.Equal(t, 1.0, UpUnit2f.Y)
	assert.Equal(t, 0.0, UpUnit2f.X)
	assert.Equal(t, -1.0, DownUnit2f.Y)
	assert.Equal(t, 0.0, DownUnit2f.X)
	assert.Equal(t, 0.0, RightUnit2f.Y)
	assert.Equal(t, 1.0, RightUnit2f.X)
	assert.Equal(t, 0.0, LeftUnit2f.Y)
	assert.Equal(t, -1.0, LeftUnit2f.X)
}

func TestNewVector2f(t *testing.T) {
	v := NewVector2f(1.0, 2.0)
	assert.Equal(t, 1.0, v.X)
	assert.Equal(t, 2.0, v.Y)
}

func TestZero(t *testing.T) {
	zero := Zero2f()
	assert.Equal(t, 0.0, zero.X)
	assert.Equal(t, 0.0, zero.Y)
}

func TestEquals(t *testing.T) {
	assert.True(t, NewVector2f(0.0, 0.0).Equals(NewVector2f(0.0, 0.0)))
	assert.True(t, NewVector2f(0.1, 0.1).Equals(NewVector2f(0.1, 0.1)))
	assert.False(t, NewVector2f(0.0, 0.0).Equals(NewVector2f(0.0, 0.1)))
	assert.False(t, NewVector2f(0.1, 0.0).Equals(NewVector2f(0.0, 0.0)))
}

func TestEqualsTo(t *testing.T) {
	assert.True(t, NewVector2f(0.0, 0.0).EqualTo(NewVector2f(0.0, 0.0), 0))
	assert.True(t, NewVector2f(0.1, 0.1).EqualTo(NewVector2f(0.1, 0.1), 0))

	assert.False(t, NewVector2f(0.001, 0.001).EqualTo(NewVector2f(0.002, 0.002), 0))
	assert.True(t, NewVector2f(0.001, 0.001).EqualTo(NewVector2f(0.002, 0.002), 0.01))
	assert.False(t, NewVector2f(-0.001, -0.001).EqualTo(NewVector2f(-0.002, -0.002), 0))
	assert.True(t, NewVector2f(-0.001, -0.001).EqualTo(NewVector2f(-0.002, -0.002), 0.01))
}

func TestSet(t *testing.T) {
	v := NewVector2f(10.0, 20.0)
	v.Set(100.0, 200.0)

	assert.Equal(t, 100.0, v.X)
	assert.Equal(t, 200.0, v.Y)
}

func TestSetX(t *testing.T) {
	v := NewVector2f(10.0, 20.0)
	v.SetX(100.0)

	assert.Equal(t, 100.0, v.X)
	assert.Equal(t, 20.0, v.Y)
}

func TestSetY(t *testing.T) {
	v := NewVector2f(10.0, 20.0)
	v.SetY(200.0)

	assert.Equal(t, 10.0, v.X)
	assert.Equal(t, 200.0, v.Y)
}

func TestAddition(t *testing.T) {
	vectors := []*Vector2f{
		NewVector2f(0.0, 0.0),
		NewVector2f(0.0, 0.0),
		NewVector2f(-2.0, -4.0),
	}

	operands := []*Vector2f{
		NewVector2f(0.0, 0.0),
		NewVector2f(1.0, 2.0),
		NewVector2f(4.0, 8.0),
	}

	expectations := [][]float64{
		{0.0, 0.0},
		{1.0, 2.0},
		{2.0, 4.0},
	}

	for i := 0; i < len(vectors); i++ {
		// immutable Add
		newVec := Add2f(vectors[i], operands[i])
		assert.Equal(t, expectations[i][0], newVec.X)
		assert.Equal(t, expectations[i][1], newVec.Y)

		// mutate with Add
		vectors[i].Add(operands[i])
		assert.Equal(t, expectations[i][0], vectors[i].X)
		assert.Equal(t, expectations[i][1], vectors[i].Y)
	}
}

func TestSubtraction(t *testing.T) {
	vectors := []*Vector2f{
		NewVector2f(0.0, 0.0),
		NewVector2f(0.0, 0.0),
		NewVector2f(-2.0, -4.0),
		NewVector2f(-2.0, -4.0),
	}

	operands := []*Vector2f{
		NewVector2f(0.0, 0.0),
		NewVector2f(1.0, 2.0),
		NewVector2f(4.0, 8.0),
		NewVector2f(-4.0, -6.0),
	}

	expectations := [][]float64{
		{0.0, 0.0},
		{-1.0, -2.0},
		{-6.0, -12.0},
		{2.0, 2.0},
	}

	for i := 0; i < len(vectors); i++ {
		// immutable Subtract
		newVec := Subtract2f(vectors[i], operands[i])
		assert.Equal(t, expectations[i][0], newVec.X)
		assert.Equal(t, expectations[i][1], newVec.Y)

		// mutate with Subtract
		vectors[i].Subtract(operands[i])
		assert.Equal(t, expectations[i][0], vectors[i].X)
		assert.Equal(t, expectations[i][1], vectors[i].Y)
	}
}

func TestMultiplication(t *testing.T) {
	vectors := []*Vector2f{
		NewVector2f(0.0, 0.0),
		NewVector2f(1.0, 1.0),
		NewVector2f(-2.0, -4.0),
		NewVector2f(-2.0, -4.0),
	}

	operands := []*Vector2f{
		NewVector2f(0.0, 0.0),
		NewVector2f(1.0, 2.0),
		NewVector2f(4.0, 8.0),
		NewVector2f(-4.0, -6.0),
	}

	expectations := [][]float64{
		{0.0, 0.0},
		{1.0, 2.0},
		{-8.0, -32.0},
		{8.0, 24.0},
	}

	for i := 0; i < len(vectors); i++ {
		// immutable Multiply
		newVec := Multiply2f(vectors[i], operands[i])
		assert.Equal(t, expectations[i][0], newVec.X)
		assert.Equal(t, expectations[i][1], newVec.Y)

		// mutate with Multiply
		vectors[i].Multiply(operands[i])
		assert.Equal(t, expectations[i][0], vectors[i].X)
		assert.Equal(t, expectations[i][1], vectors[i].Y)
	}
}

func TestDivide(t *testing.T) {
	vectors := []*Vector2f{
		NewVector2f(0.0, 0.0),
		NewVector2f(1.0, 1.0),
		NewVector2f(-2.0, -4.0),
		NewVector2f(-2.0, -4.0),
	}

	operands := []*Vector2f{
		NewVector2f(1.0, 1.0),
		NewVector2f(1.0, 2.0),
		NewVector2f(4.0, 8.0),
		NewVector2f(-4.0, -4.0),
	}

	expectations := [][]float64{
		{0.0, 0.0},
		{1.0, 0.5},
		{-0.5, -0.5},
		{0.5, 1.0},
	}

	for i := 0; i < len(vectors); i++ {
		// immutable Divide
		newVec := Divide2f(vectors[i], operands[i])
		assert.Equal(t, expectations[i][0], newVec.X)
		assert.Equal(t, expectations[i][1], newVec.Y)

		// mutate with Divide
		vectors[i].Divide(operands[i])
		assert.Equal(t, expectations[i][0], vectors[i].X)
		assert.Equal(t, expectations[i][1], vectors[i].Y)
	}
}

func TestScale(t *testing.T) {
	vectors := []*Vector2f{
		NewVector2f(0.0, 0.0),
		NewVector2f(1.0, 1.0),
		NewVector2f(-2.0, -4.0),
		NewVector2f(-2.0, -4.0),
	}

	operands := []float64{
		1.0,
		1.0,
		2.0,
		-2.0,
	}

	expectations := [][]float64{
		{0.0, 0.0},
		{1.0, 1.0},
		{-4.0, -8.0},
		{4.0, 8.0},
	}

	for i := 0; i < len(vectors); i++ {
		// immutable Scale
		newVec := Scale2f(vectors[i], operands[i])
		assert.Equal(t, expectations[i][0], newVec.X)
		assert.Equal(t, expectations[i][1], newVec.Y)

		// mutate with Multiply
		vectors[i].Scale(operands[i])
		assert.Equal(t, expectations[i][0], vectors[i].X)
		assert.Equal(t, expectations[i][1], vectors[i].Y)
	}
}

func TestMagnitude(t *testing.T) {
	vectors := []*Vector2f{
		NewVector2f(0.0, 0.0),
		NewVector2f(1.0, 1.0),
		NewVector2f(-2.0, -4.0),
		NewVector2f(-2.0, 5.0),
	}

	expectations := []float64{
		0.0,
		1.41421356237,
		4.47213595499,
		5.38516480713,
	}

	for i := 0; i < len(vectors); i++ {
		assert.GreaterOrEqual(t, EqualityEpsilon, math.Abs(expectations[i]-vectors[i].Magnitude()))
	}
}

func TestNormalize(t *testing.T) {
	vectors := []*Vector2f{
		NewVector2f(2.0, 2.0),
		NewVector2f(3.0, 12.0),
		NewVector2f(-7.0, 9),
		NewVector2f(-7.0, -7.0),
	}

	expectations := [][]float64{
		{1.0 / math.Sqrt(2), 1.0 / math.Sqrt(2)},
		{1.0 / math.Sqrt(17), 4.0 / math.Sqrt(17)},
		{-7.0 / math.Sqrt(130), 9.0 / math.Sqrt(130)},
		{-1.0 / math.Sqrt(2), -1.0 / math.Sqrt(2)},
	}

	for i := 0; i < len(vectors); i++ {
		// immutable Normalize
		newVec := Normalize2f(vectors[i])
		assert.GreaterOrEqual(t, EqualityEpsilon, math.Abs(expectations[i][0]-newVec.X))
		assert.GreaterOrEqual(t, EqualityEpsilon, math.Abs(expectations[i][1]-newVec.Y))

		// mutate with Normalize
		vectors[i].Normalize()
		assert.GreaterOrEqual(t, EqualityEpsilon, math.Abs(expectations[i][0]-vectors[i].X))
		assert.GreaterOrEqual(t, EqualityEpsilon, math.Abs(expectations[i][1]-vectors[i].Y))
	}
}

func TestDotProduct(t *testing.T) {
	vectors := [][]*Vector2f{
		{NewVector2f(-6.0, 8.0), NewVector2f(5.0, 12.0)},
		{NewVector2f(-12.0, 16.0), NewVector2f(12.0, 9.0)},
		{NewVector2f(-4.0, -9.0), NewVector2f(-1.0, 2.0)},
	}

	expectations := []float64{
		66.0,
		0.0,
		-14,
	}

	for i := 0; i < len(vectors); i++ {
		result := DotProduct2f(vectors[i][0], vectors[i][1])
		assert.GreaterOrEqual(t, EqualityEpsilon, math.Abs(expectations[i]-result))
	}
}

func TestDeterminant(t *testing.T) {
	vectors := [][]*Vector2f{
		{NewVector2f(-6.0, 8.0), NewVector2f(5.0, 12.0)},
		{NewVector2f(-12.0, 16.0), NewVector2f(12.0, 9.0)},
		{NewVector2f(-4.0, -9.0), NewVector2f(-1.0, 2.0)},
	}

	expectations := []float64{
		-32,
		84,
		1,
	}

	for i := 0; i < len(vectors); i++ {
		result := Determinant(vectors[i][0], vectors[i][1])
		assert.GreaterOrEqual(t, EqualityEpsilon, math.Abs(expectations[i]-result))
	}
}

func TestAngleBetween(t *testing.T) {
	vectors := [][]*Vector2f{
		{UpUnit2f, UpUnit2f},
		{UpUnit2f, RightUnit2f},
		{UpUnit2f, DownUnit2f},
		{UpUnit2f, LeftUnit2f},
		{UpUnit2f, NewVector2f(2.0, 2.0)},
	}

	expectations := []float64{
		0.0,
		math.Pi / 2,
		math.Pi,
		-math.Pi / 2,
		math.Pi / 4,
	}

	for i := 0; i < len(vectors); i++ {
		result := AngleBetween(vectors[i][0], vectors[i][1])
		assert.GreaterOrEqual(t, EqualityEpsilon, math.Abs(expectations[i]-result))
	}
}
