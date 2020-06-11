package vec

import (
	"testing"
)

func TestAddIncompatibleType(t *testing.T) {
	v := Vec2{
		X: 0,
		Y: 0,
	}
	s := "foo"
	_, err := v.Add(s)
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
	expected := Vec2{X: 1, Y: -1}
	res, err := v.Add(o)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	if !Equals(res, expected) {
		t.Errorf("Expected %v and got %v!", expected, res)
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

	_, err := v.Add(i0)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	_, err = v.Add(i1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	e := Vec2{1, 1}
	res, err := v.Add(i2)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
		return
	}
	if res.X != e.X && res.Y != e.Y {
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

	_, err := v.Add(f1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	res, err := v.Add(f2)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
		return
	}
	e := Vec2{1, 1}
	if res.X != e.X && res.Y != e.Y {
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

func TestMulS(t *testing.T) {
	v := Vec2{X: 1.0, Y: -2.3}
	s := 2.4
	v = v.MulS(s)
	e := Vec2{X: 2.4, Y: -5.52}
	if !Equals(v, e) {
		t.Errorf("Expected %v and got %v!", e, v)
	}
}

func TestVec2_Normalized(t *testing.T) {
	values := []Vec2{
		{-15.5,64},{0,0},{88.1,-121.1},
	}
	expected := []Vec2{
		{-0.23538270083995802, 0.9719027647585363}, {0,0},{0.5882908040762137, -0.8086494480548181},
	}
	for i,v := range values {
		n := v.Normalized()
		if !Equals(n, expected[i]) {
			t.Errorf("Expected %v and got %v!", expected[i], n)
		}
	}
}

func TestVec2_Normalize(t *testing.T) {
	values := []Vec2{
		{-15.5,64},{0,0},{88.1,-121.1},
	}
	expected := []Vec2{
		{-0.23538270083995802, 0.9719027647585363}, {0,0},{0.5882908040762137, -0.8086494480548181},
	}
	for i,v := range values {
		v.Normalize()
		if !Equals(v, expected[i]) {
			t.Errorf("Expected %v and got %v!", expected[i], v)
		}
	}
}

// TODO: check for validity
// Test chaining multiple operations together
func TestVec2_Chaining(t *testing.T) {
	val := Vec2{0.4, 2.12}.MulV(Vec2{2.5, -2.3}).Scale(0.5).Normalized().AddV(Vec2{-1.0, 2.0})
	expected := Vec2{-0.799095403593255, 1.0203891879207116}
	if !Equals(val, expected) {
		t.Errorf("Expected %v and got %v!", expected, val)
	}
}
