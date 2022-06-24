package main

// Unit - testing for comma
// tcs -- stands for Test Case Struct
// tc  -- stands for Test Case
// ret -- stands for return

// _ = Blank Identifier

// In a for loop, we can use the blank identifier
// to ignore any value returned from the range keyword. In this case, I want to ignore the index, which is the first argument returned.

import (
	"bytes"
	"os"
	"testing"
)

func TestComma(t *testing.T) {
	var tcs = []struct {
		args    []string
		expects string
	}{
		{[]string{"comma", "1", "12", "123", "1234", "1234567890"}, "  1\n  12\n  123\n  1,234\n  1,234,567,890\n"},
		{[]string{"comma", "1234567"}, "  1,234,567\n"},
	}

	for _, tc := range tcs {
		os.Args = tc.args
		stdout = new(bytes.Buffer) // captured output
		main()
		ret := stdout.(*bytes.Buffer).String()
		if ret != tc.expects {
			t.Errorf("Failed comma. Value: %v, Expects %q, Result = %q", tc.args, tc.expects, ret)
		}
	}
}
