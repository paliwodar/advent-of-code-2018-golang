package aoc2018

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestCalibrate(t *testing.T) {
	total := Calibrate(readInts("input1.txt"))
	if total != 574 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 574)
	}
}

func TestCalibrate2(t *testing.T) {
	total := Calibrate2(readInts("input1.txt"))
	if total != 452 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 452)
	}
}

func readInts(filename string) []int {
	return mapToInts(readStrings(filename))
}

func readStrings(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	return strings.Split(string(data), "\n")
}

func mapToInts(splitted []string) []int {
	integers := make([]int, len(splitted))
	for i, v := range splitted {
		t, _ := strconv.Atoi(v)
		integers[i] = t
	}
	return integers
}
