package day_02

import (
	"testing"
)

func TestExtractData(t *testing.T) {
	filePath := "input_test.txt"

	data := extractData(&filePath)

	if len(data) != 6 {
		t.Fatalf(`extractData("input_test.txt") :: report cound does not match`)
	}
}

func TestIsSorted(t *testing.T) {
	isSortedAsec := [...]int{1, 2, 2, 3, 5}
	isSortedDesc := [...]int{5, 3, 2, 1}

	unhappyPath := [...]int{5, 6, 2, 1}

	if !IsSorted(isSortedAsec[:]) {
		t.Fatalf("should mark as correctly sorted asec: %d\n", isSortedAsec)
	}

	if !IsSorted(isSortedDesc[:]) {
		t.Fatalf("should mark as correctly sorted desc: %d\n", isSortedDesc)
	}

	if IsSorted(unhappyPath[:]) {
		t.Fatalf("should mark as not sorted: %d\n", unhappyPath)
	}
}

func TestIsSafe(t *testing.T) {
	safeOne := [...]int{7, 6, 4, 2, 1}
	unsafeGreaterThan3One := [...]int{1, 2, 7, 8, 9}
	unsafeGreaterThan3Two := [...]int{9, 7, 6, 2, 1}
	unsafeChangeDirectionOne := [...]int{1, 3, 2, 4, 5}
	unsafeNoChangeOne := [...]int{8, 6, 4, 4, 1}
	safeTwo := [...]int{1, 3, 6, 7, 9}

	if !IsSafe(safeOne[:]) {
		t.Fatalf("should be safe One: %d\n", safeOne)
	}

	if !IsSafe(safeTwo[:]) {
		t.Fatalf("should be safe Two: %d\n", safeTwo)
	}

	if IsSafe(unsafeGreaterThan3One[:]) {
		t.Fatalf("should be unsafe, increase of 5: %d\n", unsafeGreaterThan3One)
	}

	if IsSafe(unsafeGreaterThan3Two[:]) {
		t.Fatalf("should be unsafe, decrease of 4: %d\n", unsafeGreaterThan3Two)
	}

	if IsSafe(unsafeChangeDirectionOne[:]) {
		t.Fatalf("should be unsafe, direction changed: %d\n", unsafeChangeDirectionOne)
	}

	if IsSafe(unsafeNoChangeOne[:]) {
		t.Fatalf("should be unsafe, neither an increase or a decrease: %d\n", unsafeNoChangeOne)
	}
}

func TestProblemOneTestData(t *testing.T) {
	filePath := "input_test.txt"
	data := extractData(&filePath)
	total := ProblemOne(data)

	if total != 2 {
		t.Fatalf("WRONG - expected: %d, actual: %d/n", 2, total)
	}
	t.Logf("First Question: %d\n", total)
}

func TestProblemOne(t *testing.T) {
	filePath := "input.txt"
	data := extractData(&filePath)
	total := ProblemOne(data)

	if total != 306 {
		t.Fatalf("WRONG - expected: %d, actual: %d/n", 306, total)
	}
	t.Logf("First Question: %d\n", total)
}

func TestProblemTwoTestData(t *testing.T) {
	filePath := "input_test.txt"
	data := extractData(&filePath)
	total := ProblemTwo(data)

	if total != 4 {
		t.Fatalf("WRONG - expected: %d, actual: %d/n", 4, total)
	}
	t.Logf("First Question: %d\n", total)
}

func TestProblemTwo(t *testing.T) {
	filePath := "input.txt"
	data := extractData(&filePath)
	total := ProblemTwo(data)

	// First attempt gave 627. AOC said this is too high. adding to test
	if total >= 627 {
		t.Fatalf("TO HIGH: %d\n", total)
	}

	if total <= 346 {
		t.Fatalf("TO LOW: %d\n", total)
	}

	t.Logf("Second Question: %d\n", total)
}
