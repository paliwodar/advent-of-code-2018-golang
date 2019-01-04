package aoc2018

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestCalibrate(t *testing.T) {
	total := Calibrate(readChanges())
	if total != 574 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 574)
	}
}

func TestCalibrate2(t *testing.T) {
	total := Calibrate2(readChanges())
	if total != 452 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 452)
	}
}

func readChanges() []int {
	data, _ := ioutil.ReadFile("input1.txt")
	splitted := strings.Split(string(data), "\n")
	integers := make([]int, len(splitted))

	for i, v := range splitted {
		t, _ := strconv.Atoi(v)
		integers[i] = t
	}
	return integers
}
