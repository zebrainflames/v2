package v2

import (
	"testing"
)

func TestAddIncompatibleType(t *testing.T) {
	v := &Vec2{
		X: 0,
		Y: 0,
	}
	s := "foo"
	err := v.Add(s)
	if err == nil {
		// TODO: what's a good test for this?
		t.Fatal("Expected string input to Add to fail!")
	}

}

func TestAddGenericV2(t *testing.T) {
	v := &Vec2{
		X: 0,
		Y: 0,
	}
	o := &Vec2{
		X: 1,
		Y: -1,
	}
	err := v.Add(o)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
}

func TestAddGenericInts(t *testing.T) {
	var i0 = 1
	var i1 int32 = 1
	var i2 int64 = 1

	v := &Vec2{
		X: 0,
		Y: 0,
	}

	err := v.Add(i0)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	err = v.Add(i1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	err = v.Add(i2)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	e := &Vec2{3, 3}
	if v.X != e.X && v.Y != e.Y {
		t.Error("Expected ", e, ", got ", v)
	}
}

func TestAddGenericFloats(t *testing.T) {
	var f1 float32 = 1
	var f2 float64 = 1

	v := &Vec2{
		X: 0,
		Y: 0,
	}

	err := v.Add(f1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	err = v.Add(f2)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	e := &Vec2{2, 2}
	if v.X != e.X && v.Y != e.Y {
		t.Error("Expected ", e, ", got ", v)
	}
}

func TestEqualsTrue(t *testing.T) {
	v := Vec2{1, -23}
	o := Vec2{12, -123}
	if Equals(v, o) != false {
		t.Errorf("Expected %v and %v to not be equal!", v, o)
	}
}

