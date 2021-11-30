package day01

import (
	"testing"
)

func Test_Day01_One(t *testing.T) {
	ans := run_one()
	expected := 1
	if ans != expected {
		t.Errorf("Expected %d; was %d", expected, ans)
	}
}

func Test_Day01_Two(t *testing.T) {
	ans := run_two()
	expected := 2
	if ans != expected {
		t.Errorf("Expected %d; was %d", expected, ans)
	}
}
