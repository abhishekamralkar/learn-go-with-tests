package main

func TesSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	wanted := 15

	if got != wanted {
		t.Errorf("got %d want %d given %v", got, wanted, numbers)
	}
}
