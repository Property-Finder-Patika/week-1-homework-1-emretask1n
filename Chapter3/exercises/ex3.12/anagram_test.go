package main

// Unit - testing for anagram
// tcs -- stands for Test Case Struct
// tc  -- stands for Test Case
// ret -- stands for return

// _ = Blank Identifier

// In a for loop, we can use the blank identifier
// to ignore any value returned from the range keyword. In this case, I want to ignore the index, which is the first argument returned.

import (
	"reflect"
	"runtime"
	"testing"
)

func testIsAnagram(t *testing.T, f func(string, string) bool) {
	tcs := []struct {
		s1      string
		s2      string
		expects bool
	}{
		{"abc", "cba", true},
		{"abc", "abc", false},
		{"abc", "abcd", false},
		{"abc", "ab", false},
		{"abc", "", false},
	}

	for _, tc := range tcs {
		ret := f(tc.s1, tc.s2)
		if ret != tc.expects {
			t.Errorf("Failed %v, s1: %s, s2: %s, result: %v", getFunctionName(f), tc.s1, tc.s2, ret)
		}
	}
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func TestIsAnagramReflect(t *testing.T) {
	testIsAnagram(t, isAnagramReflect)
}

func TestIsAnagramMap(t *testing.T) {
	testIsAnagram(t, isAnagramMap)
}

func TestIsAnagramRemove(t *testing.T) {
	testIsAnagram(t, isAnagramRemove)
}
