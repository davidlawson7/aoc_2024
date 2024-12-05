package main

import (
	"testing"
)

func TestCountAllWordsTestData(t *testing.T) {

	w := WordSearchInit("input_test.txt")
	wc := w.FindAllXMAS()

	if wc != 18 {
		t.Fatalf("[FindAllXMAS] should have a total of 18 words, actual: %d\n", wc)
	}
}
