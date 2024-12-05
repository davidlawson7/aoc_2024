package main

import (
	"testing"
)

func TestCountAllWordsTestData(t *testing.T) {
	w := WordSearchInit("input_test.txt")
	wc := w.CheckEachByte()
	if wc != 18 {
		t.Fatalf("[CountAllWords] should have a total of 18 words, actual: %d\n", wc)
	}
	t.Logf("Word Count: %d\n", wc)
}
