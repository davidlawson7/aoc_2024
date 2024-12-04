package day_04

import (
	"testing"
)

func TestGetInput(t *testing.T) {
	testFileName := "input_test.txt"
	wordSearch := GetInput(&testFileName)
	if len((*wordSearch)) != 10 {
		t.Fatalf("[GetInput] should have a total of 10 rows, actual: %d\n", len((*wordSearch)))
	}

	if len((*wordSearch)[0]) != 10 {
		t.Fatalf("[GetInput] should have a total of 10 columns, actual: %d\n", len((*wordSearch)[0]))
	}
}

func TestCountAllWordsTestData(t *testing.T) {
	testFileName := "input_test.txt"
	wc := CountAllWords(&testFileName)

	if wc != 18 {
		t.Fatalf("[CountAllWords] should have a total of 18 words, actual: %d\n", wc)
	}
	t.Logf("Word Count: %d\n", wc)
}

func TestCountAllWords(t *testing.T) {
	testFileName := "input.txt"
	wc := CountAllWords(&testFileName)

	if wc >= 2483 {
		t.Fatalf("[CountAllWords] Too high: %d\n", wc)
	}

	t.Logf("Word Count: %d\n", wc)
}
