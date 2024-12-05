package main

import (
	"testing"
)

func TestCountAllWordsTestData(t *testing.T) {

	w := WordSearchInit("input_test.txt")
	wc, _ := w.FindAllXMAS()

	if wc != 18 {
		t.Fatalf("[FindAllXMAS] should have a total of 18 words, actual: %d\n", wc)
	}
}

func TestCountAllXMASTestData(t *testing.T) {

	w := WordSearchInit("input_test.txt")
	_, xm := w.FindAllXMAS()

	if xm != 9 {
		t.Fatalf("[FindAllXMAS] should have a total of 9 words, actual: %d\n", xm)
	}
}
