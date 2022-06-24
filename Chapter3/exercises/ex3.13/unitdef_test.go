package main

import "testing"

func TestUnitDef(t *testing.T) {
	if step() != 1000 {
		t.Error()
	}
}
