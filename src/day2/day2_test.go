package day2

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	input := [...]string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	res := partOne(input[:])
	if res != 150 {
		t.Error("Expected 150")
	}
}

func TestPartTwo(t *testing.T) {
	input := [...]string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	res := partTwo(input[:])
	if res != 900 {
		t.Error("Expected 900")
	}
}
