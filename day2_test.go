package aoc2018

import (
	"testing"
)

func TestCalculateChecksum(t *testing.T) {
	total := CalculateChecksum(readStrings("input2.txt"))
	if total != 7192 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 7192)
	}
}

func TestCalculateChecksum2(t *testing.T) {
	total := CalculateChecksum2(readStrings("input2.txt"))
	if total != "mbruvapghxlzycbhmfqjonsie" {
		t.Errorf("Sum was incorrect, got: %s, want: %s.", total, "mbruvapghxlzycbhmfqjonsie")
	}
}

