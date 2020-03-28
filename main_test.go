package main

import "testing"

func TestMultiply(t *testing.T) {
	res := Multiply(11)
	if res != 121 {
		t.Error()
	}
}
