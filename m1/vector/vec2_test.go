package vector

import (
	"fmt"
	"testing"
)

func TestAddVec2(t *testing.T) {

	var tests = []struct {
		in1      Vec2
		in2      Vec2
		expected Vec2
	}{
		{Vec2{}, Vec2{}, Vec2{}},
		{Vec2{1, 2}, Vec2{2, 2}, Vec2{3, 4}},
	}

	for num, test := range tests {
		testname := fmt.Sprintf("T=%v", num)
		t.Run(testname, func(t *testing.T) {
			got := AddVec2(test.in1, test.in2)
			if got.X != test.expected.X || got.Y != test.expected.Y {
				t.Errorf("\nCalled:AddVec2(%v,%v)\nExpected:%v\nGot:%v", test.in1, test.in2, test.expected, got)
			}
		})
	}
}

func TestNegate1(t *testing.T) {

	v1 := Vec2{1, 1}
	v1.Negate()

	if v1.X != -1 || v1.Y != -1 {
		t.Errorf("\nCalled:Negate(1,1)\nExpected:Vec2(-1,-1)\nGot:%v", v1)
	}
}
