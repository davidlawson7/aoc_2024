package day_04

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type WSBoundsError struct {
	row     int
	col     int
	message string
}

func (e *WSBoundsError) Error() string {
	return fmt.Sprintf("%d, %d - %s", e.row, e.col, e.message)
}

type WordSearch struct {
	Grid *[][]byte
}

func (w WordSearch) GetLetter(row int, col int) (letter byte, err error) {
	// Do some out of bounds errorings... eyeroll
	mRow, mCol := w.GetBounds()

	if row > mRow || row < 0 || col > mCol || col < 0 {
		return 'a', &WSBoundsError{row, col, "can't work with it"}
	}

	return (*w.Grid)[row][col], nil
}

func (w WordSearch) GetBounds() (row int, col int) {
	return len((*w.Grid)) - 1, len((*w.Grid)[0]) - 1
}

func (w WordSearch) IsX(row int, col int) bool {
	if letter, err := w.GetLetter(row, col); err == nil {
		return letter == 'X'
	}
	return false // Should return a error as well if we get a out of bounds issue... lazy
}

func (w WordSearch) CheckHorizontal(row int, col int) bool {
	if m, err := w.GetLetter(row, col+1); err == nil && m == 'M' {
		if a, err := w.GetLetter(row, col+2); err == nil && a == 'A' {
			if s, err := w.GetLetter(row, col+3); err == nil && s == 'S' {
				return true
			}
		}
	}
	return false
}

func (w WordSearch) CheckBackwards(row int, col int) bool {
	if m, err := w.GetLetter(row, col-1); err == nil && m == 'M' {
		if a, err := w.GetLetter(row, col-2); err == nil && a == 'A' {
			if s, err := w.GetLetter(row, col-3); err == nil && s == 'S' {
				return true
			}
		}
	}
	return false
}

func (w WordSearch) CheckAbove(row int, col int) bool {
	if m, err := w.GetLetter(row-1, col); err == nil && m == 'M' {
		if a, err := w.GetLetter(row-2, col); err == nil && a == 'A' {
			if s, err := w.GetLetter(row-3, col); err == nil && s == 'S' {
				return true
			}
		}
	}
	return false
}

func (w WordSearch) CheckBelow(row int, col int) bool {
	if m, err := w.GetLetter(row+1, col); err == nil && m == 'M' {
		if a, err := w.GetLetter(row+2, col); err == nil && a == 'A' {
			if s, err := w.GetLetter(row+3, col); err == nil && s == 'S' {
				return true
			}
		}
	}
	return false
}

func (w WordSearch) CheckTopLeft(row int, col int) bool {
	if m, err := w.GetLetter(row-1, col-1); err == nil && m == 'M' {
		if a, err := w.GetLetter(row-2, col-2); err == nil && a == 'A' {
			if s, err := w.GetLetter(row-3, col-3); err == nil && s == 'S' {
				return true
			}
		}
	}
	return false
}

func (w WordSearch) CheckTopRight(row int, col int) bool {
	if m, err := w.GetLetter(row-1, col+1); err == nil && m == 'M' {
		if a, err := w.GetLetter(row-2, col+2); err == nil && a == 'A' {
			if s, err := w.GetLetter(row-3, col+3); err == nil && s == 'S' {
				return true
			}
		}
	}
	return false
}

func (w WordSearch) CheckBottomLeft(row int, col int) bool {
	if m, err := w.GetLetter(row+1, col-1); err == nil && m == 'M' {
		if a, err := w.GetLetter(row+2, col-2); err == nil && a == 'A' {
			if s, err := w.GetLetter(row+3, col-3); err == nil && s == 'S' {
				return true
			}
		}
	}
	return false
}

func (w WordSearch) CheckBottomRight(row int, col int) bool {
	if m, err := w.GetLetter(row+1, col+1); err == nil && m == 'M' {
		if a, err := w.GetLetter(row+2, col+2); err == nil && a == 'A' {
			if s, err := w.GetLetter(row+3, col+3); err == nil && s == 'S' {
				return true
			}
		}
	}
	return false
}

// Transforms raw input from a text file into something workable for this test.
func GetInput(filePath *string) *[][]byte {
	file, err := os.Open(*filePath)

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var wordSearch [][]byte

	for fileScanner.Scan() {
		wordSearch = append(wordSearch, fileScanner.Bytes())
	}

	file.Close()

	return &wordSearch
}

// Checks for the word XMAS starting from (aka letter X) the given position
// (x & y) within the wordSearch.
func Check(w *WordSearch, row int, col int) int {
	if !w.IsX(row, col) {
		return 0
	}

	xmasCount := 0

	if w.CheckBackwards(row, col) {
		xmasCount += 1
	}

	if w.CheckHorizontal(row, col) {
		xmasCount += 1
	}

	if w.CheckAbove(row, col) {
		xmasCount += 1
	}

	if w.CheckBelow(row, col) {
		xmasCount += 1
	}

	if w.CheckTopLeft(row, col) {
		xmasCount += 1
	}

	if w.CheckTopRight(row, col) {
		xmasCount += 1
	}

	if w.CheckBottomLeft(row, col) {
		xmasCount += 1
	}

	if w.CheckBottomRight(row, col) {
		xmasCount += 1
	}

	return xmasCount
}

func CountAllWords(filePath *string) int {
	wordSearch := GetInput(filePath)
	w := WordSearch{
		Grid: wordSearch,
	}

	r, c := w.GetBounds()
	sum := 0

	for i := 0; i < r+1; i++ {
		for j := 0; j < c+1; j++ {
			sum = sum + Check(&w, i, j)
		}
	}
	return sum
}
