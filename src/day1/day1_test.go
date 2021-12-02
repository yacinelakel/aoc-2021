package day1

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	input := [...]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	res := partOne(input[:])

	if res != 7 {
		t.Error("Expected 7")
	}
}

func TestPartTwo(t *testing.T) {
	input := [...]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	res := partTwo(input[:])

	if res != 5 {
		t.Error("Expected 5")
	}
}
