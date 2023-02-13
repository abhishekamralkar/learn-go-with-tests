package main

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{11, 12, 13, 14, 15}

	got := Sum(numbers)
	want := 56

	if got != want {
		t.Errorf("%d got - %d want - %v given\n", got, want, given)
	}

}
