package iterations

import "testing"

func TestLooper(t *testing.T) {
	got := Looper("a")
	expected := "aaaaa"

	if got != expected {
		t.Errorf("got %s expected %s\n", got, expected)
	}
}
