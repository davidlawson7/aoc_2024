package main

import (
	"slices"
	"testing"
)

func TestExtractData(t *testing.T) {
	filePath := "input_test.txt"
	expectedLeft := [6]int{3, 4, 2, 1, 3, 3}
	expectedRight := [6]int{4, 3, 5, 3, 9, 3}

	left, right := extractData(&filePath)

	if !slices.Equal(*left, expectedLeft[:]) {
		t.Fatalf(`extractData("input_test.txt") :: left array does not match`)
	}

	if !slices.Equal(*right, expectedRight[:]) {
		t.Fatalf(`extractData("input_test.txt") :: right array does not match`)
	}
}

func TestCalcDistances(t *testing.T) {
	// 1, 2, 3, 3, 3, 4
	left := [6]int{3, 4, 2, 1, 3, 3}
	// 3, 3, 3, 4, 5, 9
	right := [6]int{4, 3, 5, 3, 9, 3}

	expected := [6]int{2, 1, 0, 1, 2, 5}
	actual := calcDistances(left[:], right[:])

	if !slices.Equal(*actual, expected[:]) {
		t.Fatalf(`calcDistances(left, right) :: expected: %d, actual: %d`, expected, *actual)
	}
}

func TestTotalDistance(t *testing.T) {
	distances := [6]int{2, 1, 0, 1, 2, 5}
	expected := 11
	actual := totalDistance(distances[:])

	if actual != expected {
		t.Fatalf(`totalDistance(distances) :: expected: %d, actual: %d`, expected, actual)
	}
}

func TestFindScores(t *testing.T) {
	// 1, 2, 3, 3, 3, 4
	left := [6]int{3, 4, 2, 1, 3, 3}
	leftS := left[:]
	leftP := &leftS
	// 3, 3, 3, 4, 5, 9
	right := [6]int{4, 3, 5, 3, 9, 3}
	rightS := right[:]
	rightP := &rightS

	expected := [6]int{9, 4, 0, 0, 9, 9}
	actual := findScores(leftP, rightP)

	if !slices.Equal(*actual, expected[:]) {
		t.Fatalf(`calcDistances(left, right) :: expected: %d, actual: %d`, expected, *actual)
	}
}
