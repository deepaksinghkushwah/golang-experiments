package main

import "testing"

func TestMySum(t *testing.T) {
	sum := mySum(1, 2, 3)

	if sum != 6 {
		t.Error("Expacting 6 got", sum)
	}
}
