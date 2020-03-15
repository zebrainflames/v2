package v2

import (
	"testing"
)

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
	if v.X != 3 && v.Y != 3 {
		t.Error("Expected ", e, ", got ", v)
	}
}
