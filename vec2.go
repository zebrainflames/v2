package v2

import (
	"fmt"
	"math"
)

// Vec2 is a 2D Vec2 type
type Vec2 struct {
	X,Y	float64
}

// TypeError is a type alias for errors returned from the "type agnostic" operator methods exported by this package.
type TypeError error

// Add is a more type agnostic Add shorthand, that casts some compatible types to float64. Use specific versions whenever
// possible. When a non-supported type is encountered, returns a specific error type alias TypeError.
// Currently supports Vec2, *Vec2, ints and floats
func (v *Vec2) Add(i interface{}) TypeError {
	switch o := i.(type) {
	case Vec2:
		v.AddV(o)
	case *Vec2:
		v.AddV(*o)
	case int:
		v.AddS(float64(o))
	case int32:
		v.AddS(float64(o))
	case int64:
		v.AddS(float64(o))
	case float32:
		v.AddS(float64(o))
	case float64:
		v.AddS(o)
	default:
		return fmt.Errorf("Incompatible type: %t\n", o)
	}
	return nil
}

// AddV adds vec2 v component-wise to vec2 v in-place (mutates it).
func (v *Vec2) AddV(o Vec2) {
	v.X += o.X
	v.Y += o.Y
}

// AddS adds scalar o component-wise to vec2 v in-place (mutates it).
func (v *Vec2) AddS(o float64) {
	v.X += o
	v.Y += o
}

// Add is a more type agnostic Add shorthand, that casts some compatible types to float64. Use specific versions whenever
// possible. When a non-supported type is encountered, returns a specific error type alias Vec2TypeError.
// Currently supports Vec2, *Vec2, ints and floats
func (v *Vec2) Mul(i interface{}) TypeError {
	switch o := i.(type) {
	case Vec2:
		v.MulV(o)
	case *Vec2:
		v.MulV(*o)
	case int:
		v.MulS(float64(o))
	case int32:
		v.MulS(float64(o))
	case int64:
		v.MulS(float64(o))
	case float32:
		v.MulS(float64(o))
	case float64:
		v.MulS(o)
	default:
		return fmt.Errorf("Incompatible type: %t\n", o)
	}
	return nil
}

// MulV multiplies v by vec2 o in-place (mutates it).
func (v *Vec2) MulV(o Vec2) {
	v.X *= o.X
	v.Y *= o.Y
}

// MulS multiplies v by scalar o in-place (mutates it).
func (v *Vec2) MulS(o float64) {
	v.X *= o
	v.Y *= o
}

// Sub is a type agnostic substraction shorthand. Returns specific error type alias. Use typed versions when possible.
func (v *Vec2) Sub(i interface{}) TypeError {
	switch o := i.(type) {
	case Vec2:
		v.SubV(o)
	case int:
		v.SubS(float64(o))
	case float32:
		v.SubS(float64(o))
	default:
		return fmt.Errorf("Incompatible type: %t\n", o)
	}
	return nil
}

func (v *Vec2) SubV(o Vec2) {
	v.X -= o.X
	v.Y -= o.Y
}

func (v *Vec2) SubS(o float64) {
	v.X -= o
	v.Y -= o
}

func (v *Vec2) Div(i interface{}) {
	switch o := i.(type) {
	case Vec2:
		v.DivV(o)
	case int:
		v.DivS(float64(o))
	case float32:
		v.DivS(float64(o))
	default:
		panic("Incompatible type!")
	}
}

// DivV divides v component-wise by vec2 o
func (v *Vec2) DivV(o Vec2) {
	v.X /= o.X
	v.Y /= o.Y
}

// DivS divides v component-wise by scalar o (mutates it in place)
func (v *Vec2) DivS(o float64) {
	v.X /= o
	v.Y /= o
}

// Project projects vector v to axis o
func (v *Vec2) Project(o Vec2) {
	dot := v.Dot(o)
	v.X = (dot / (o.X*o.X + o.Y*o.Y)) * o.X
	v.Y = (dot / (o.X*o.X + o.Y*o.Y)) * o.Y
}

// Dot returns the dot product for v and o
func (v Vec2) Dot(o Vec2) float64 {
	return v.X*o.X + v.Y*o.Y
}

// Cross returns the cross product for v and o
func (v Vec2) Cross(o Vec2) float64 {
	return v.X*o.Y - v.Y*o.X
}

// CrossF returns the component-wise scalar cross-product of v
func (v Vec2) CrossF(o float64) Vec2 {
	return Vec2{-v.Y * o, v.X * o}
}

// LengthSquared returns the squared length of v. Useful for distance checks etc.
func (v Vec2) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y
}

// Length returns the scalar magnitude of v. Use LengthSquared whenever possible - it's cheaper for distance checks etc.
func (v Vec2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Normalize normalizes v (scales it to unit-length 1) in-place (mutates it)
func (v *Vec2) Normalize() {
	v.MulS(1.0 / v.Length())
}

// Normalized returns a normalized (unit length of 1) version of v
func (v Vec2) Normalized() Vec2 {
	return Scale(v, 1.0/v.Length())
}

// Add adds to Vec2s together, returns the result
func Add(v, o Vec2) Vec2 {
	return Vec2{v.X + o.X, v.Y + o.Y}
}

// Sub two vectors
func Sub(v, o Vec2) Vec2 {
	return Vec2{v.X - o.X, v.Y - o.Y}
}

// Div divides v by o component-wise, returns result
func Div(v, o Vec2) Vec2 {
	return Vec2{
		X: v.X / o.X,
		Y: v.Y / o.Y,
	}
}

// Mul multiplies v by o component-wise, returns result
func Mul(v, o Vec2) Vec2 {
	return Vec2{
		X: v.X * o.X,
		Y: v.Y * o.Y,
	}
}

// Scale multiplies v by scalar r (component-wise)
func Scale(v Vec2, r float64) Vec2 {
	return Vec2{v.X * r, v.Y * r}
}

// Equals is a simple component-wise equality check
func Equals(v, o Vec2) bool {
	return v.X == o.X && v.Y == o.Y
}
