package main

import (
	"testing"
)

func TestPrintQueue(t *testing.T) {

	pq := InitPrintQueue("input_test.txt")
	p1, p2 := pq.GetCorrectUpdates()

	if p1 != 143 {
		t.Fatalf("[GetCorrectUpdates] p1 output should be 143 for test file, actual: %d\n", p1)
	}

	if p2 != 123 {
		t.Fatalf("[GetCorrectUpdates] p=2 output should be 123 for test file, actual: %d\n", p2)
	}
}
