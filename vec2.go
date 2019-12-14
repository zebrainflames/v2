package v2

import (
	"math"
)

// Vec2 is a 2D Vec2 type
type Vec2 struct {
	X,Y	float64
}


func (v *Vec2) Add(i interface{}) {
	switch o := i.(type) {
	case Vec2:
		v.AddV(o)
	case int:
		v.AddS(float64(o))
	case float32:
		v.AddS(float64(o))
	default:
		panic("Incompatible type!")
	}
}

func (v *Vec2) AddV(o Vec2) {
	v.X += o.X
	v.Y += o.Y
}

func (v *Vec2) AddS(o float64) {
	v.X += o
	v.Y += o
}

func (v *Vec2) Mul(i interface{}) {
	switch o := i.(type) {
	case Vec2:
		v.MulV(o)
	case int:
		v.MulS(float64(o))
	case float32:
		v.MulS(float64(o))
	default:
		panic("Incompatible type!")
	}
}

func (v *Vec2) MulV(o Vec2) {
	v.X *= o.X
	v.Y *= o.Y
}


func (v *Vec2) MulS(o float64) {
	v.X *= o
	v.Y *= o
}

func (v *Vec2) Sub(o Vec2) {
	v.X -= o.X
	v.Y -= o.Y
}


func (v *Vec2) Project(o Vec2) {
	dot := v.Dot(o)
	v.X = (dot / (o.X*o.X + o.Y*o.Y)) * o.X
	v.Y = (dot / (o.X*o.X + o.Y*o.Y)) * o.Y
}

func (v Vec2) Dot(o Vec2) float64 {
	return v.X*o.X + v.Y*o.Y
}

func (v Vec2) Cross(o Vec2) float64 {
	return v.X*o.Y - v.Y*o.X
}

func (v Vec2) CrossF(o float64) Vec2 {
	return Vec2{-v.Y * o, v.X * o}
}

func (v Vec2) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v Vec2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vec2) Normalize() {
	v.MulS(1.0 / v.Length())
}

func (v Vec2) Normalized() Vec2 {
	return Scale(v, 1.0/v.Length())
}

func Add(v, u Vec2) Vec2 {
	return Vec2{v.X + u.X, v.Y + u.Y}
}

func Sub(v, u Vec2) Vec2 {
	return Vec2{v.X - u.X, v.Y - u.Y}
}

func Scale(v Vec2, r float64) Vec2 {
	return Vec2{v.X * r, v.Y * r}
}
