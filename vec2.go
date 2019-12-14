package v2

import "math"

// Vec2 is a 2D Vec2 type
type Vec2 struct {
	X,Y	float64
}


func (v *Vec2) Add(o Vec2) {
	v.X += o.X
	v.Y += o.Y
}

func (v *Vec2) Sub(o Vec2) {
	v.X -= o.X
	v.Y -= o.Y
}

func (v *Vec2) Mul(o float64) {
	v.X *= o
	v.Y *= o
}

func (v Vec2) Dot(o Vec2) float64 {
	return v.X*o.X + v.Y*o.Y
}

func (v Vec2) Cross(o Vec2) float64 {
	return v.X*o.Y - v.Y*o.X
}

func (v Vec2) Crossf(o float64) Vec2 {
	return Vec2{-v.Y * o, v.X * o}
}

func (v Vec2) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v Vec2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vec2) Normalize() {
	v.Mul(1.0 / v.Length())
}

func (v Vec2) Normalized() Vec2 {
	return Scale(v, 1.0/v.Length())
}


func (v Vec2) Times(r float64) Vec2 {
	return Scale(v, r)
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
