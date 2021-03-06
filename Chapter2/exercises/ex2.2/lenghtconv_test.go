package main

import "testing"

// Unit - testing for lengthconv
// tcs -- stands for Test Case Struct
// tc  -- stands for Test Case
// ret -- stands for return

// _ = Blank Identifier

// In a for loop, we can use the blank identifier
// to ignore any value returned from the range keyword. In this case, I want to ignore the index, which is the first argument returned.

func TestFToM(t *testing.T) {
	tcs := []struct {
		length  Feet
		expects string
	}{
		{Feet(3.28), "1.00 M"},
		{Feet(1.0), "0.30 M"},
		{Feet(0.0), "0.00 M"},
	}

	for _, tc := range tcs {
		ret := FToM(tc.length)
		if ret.String() != tc.expects {
			t.Errorf("Failed FToM. Feet: %g, expects: %s, got %s", tc.length, tc.expects, ret)
		}
	}
}

func TestMToF(t *testing.T) {
	tcs := []struct {
		length  Meter
		expects string
	}{
		{Meter(1.0), "3.28 FT"},
		{Meter(0.3), "0.98 FT"},
		{Meter(0.0), "0.00 FT"},
	}

	for _, tc := range tcs {
		ret := MToF(tc.length)
		if ret.String() != tc.expects {
			t.Errorf("Failed MToF. Meter: %g, expects: %s, got %s", tc.length, tc.expects, ret)
		}
	}
}
