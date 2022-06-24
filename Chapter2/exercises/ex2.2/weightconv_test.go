package main

import "testing"

// Unit - testing for weightconv
// tcs -- stands for Test Case Struct
// tc  -- stands for Test Case
// ret -- stands for return

// _ = Blank Identifier

// In a for loop, we can use the blank identifier
// to ignore any value returned from the range keyword. In this case, I want to ignore the index, which is the first argument returned.

func TestPToKG(t *testing.T) {
	tcs := []struct {
		weight  Pound
		expects string
	}{
		{Pound(2.2046), "1.00 KG"},
		{Pound(1.0), "0.45 KG"},
		{Pound(0.0), "0.00 KG"},
	}

	for _, tc := range tcs {
		ret := PToKG(tc.weight)
		if ret.String() != tc.expects {
			t.Errorf("Failed PToKG. Pound: %g, expects: %s, got %s", tc.weight, tc.expects, ret)
		}
	}
}

func TestKGToP(t *testing.T) {
	tcs := []struct {
		weight  Kilogram
		expects string
	}{
		{Kilogram(1.0), "2.20 P"},
		{Kilogram(0.452), "1.00 P"},
		{Kilogram(0.0), "0.00 P"},
	}

	for _, tc := range tcs {
		ret := KGToP(tc.weight)
		if ret.String() != tc.expects {
			t.Errorf("Failed PToKG. Kilogram: %g, expects: %s, got %s", tc.weight, tc.expects, ret)
		}
	}
}
