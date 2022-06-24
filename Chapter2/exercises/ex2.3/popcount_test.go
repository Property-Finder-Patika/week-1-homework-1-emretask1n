package main

import "testing"

// Unit - testing for popcount

// tcs -- stands for Test Case Struct
// tc  -- stands for Test Case
// ret -- stands for return
// _ = Blank Identifier

// In a for loop, we can use the blank identifier
// to ignore any value returned from the range keyword. In this case, I want to ignore the index, which is the first argument returned.

func TestPopCountTable(t *testing.T) {
	tcs := []struct {
		number  uint64
		expects int
	}{
		{0x1234567890ABCDEF, 32},
		{0xFFFFFFFFFFFFFFFF, 64},
		{0x0000000000000002, 1},
		{0x0000000000000000, 0},
		{0x1000000000000000, 1},
	}

	for _, tc := range tcs {
		ret := PopCountTable(tc.number)
		if ret != tc.expects {
			t.Errorf("Failed PopCountTable. Number: %X, expect counts: %d, get: %d", tc.number, tc.expects, ret)
		}
	}
}

func TestPopCountLoop(t *testing.T) {
	tcs := []struct {
		number  uint64
		expects int
	}{
		{0x1234567890ABCDEF, 32},
		{0xFFFFFFFFFFFFFFFF, 64},
		{0x0000000000000002, 1},
		{0x0000000000000000, 0},
		{0x1000000000000000, 1},
	}

	for _, tc := range tcs {
		ret := PopCountLoop(tc.number)
		if ret != tc.expects {
			t.Errorf("Failed PopCountLoop. Number: %X, expect counts: %d, get: %d", tc.number, tc.expects, ret)
		}
	}
}

func BenchmarkPopCountTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountTable(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(0x1234567890ABCDEF)
	}
}
