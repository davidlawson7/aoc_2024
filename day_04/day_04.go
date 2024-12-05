package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type WordSearch struct {
	Grid       *[][]string
	Directions [][]int
}

func WordSearchInit(filePath string) *WordSearch {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var wordSearch [][]string

	for fileScanner.Scan() {
		var row []string
		for _, byte := range fileScanner.Bytes() {
			row = append(row, string(byte))
		}
		wordSearch = append(wordSearch, row)
	}

	file.Close()

	return &WordSearch{
		Grid: &wordSearch,
		Directions: [][]int{
			{0, 1},   // Left to right
			{0, -1},  // Right to left
			{-1, 0},  // Up
			{1, 0},   // Down
			{-1, 1},  // Diagonal top-left
			{-1, -1}, // Diagonal top-right
			{1, 1},   // Diagonal bottom-left
			{1, -1},  // Diagonal bottom-right
		},
	}
}

func (w WordSearch) GetBounds() (row int, col int) {
	return len((*w.Grid)) - 1, len((*w.Grid)[0]) - 1
}

func (w WordSearch) GetLetter(row int, col int) string {
	mRow, mCol := w.GetBounds()
	if (row > mRow || row < 0) || (col > mCol || col < 0) {
		return "" // Just return a empty string, will make any check fail. Could error handle if less lazy
	}
	return (*w.Grid)[row][col]
}

func (w WordSearch) CheckDirection(row int, col int, dr int, dc int) bool {
	return (w.GetLetter(row, col) + w.GetLetter(row+(1*dr), col+(1*dc)) + w.GetLetter(row+(2*dr), col+(2*dc)) + w.GetLetter(row+(3*dr), col+(3*dc))) == "XMAS"
}

func (w WordSearch) CheckEachDirection(row int, col int) int {
	count := 0
	if l := w.GetLetter(row, col); l == "X" {
		for _, d := range w.Directions {
			if w.CheckDirection(row, col, d[0], d[1]) {
				count++
			}
		}
	}
	return count
}

func (w WordSearch) CheckForMAS(row int, col int) bool {
	lt := w.GetLetter(row-1, col-1)
	rb := w.GetLetter(row+1, col+1)

	lb := w.GetLetter(row+1, col-1)
	rt := w.GetLetter(row-1, col+1)

	if ((lt == "M" && rb == "S") || (lt == "S" && rb == "M")) && ((lb == "M" && rt == "S") || (lb == "S" && rt == "M")) {
		return true
	}
	return false
}

func (w WordSearch) CheckForTwoDiagonalXMAS(row int, col int) int {
	count := 0
	if l := w.GetLetter(row, col); l == "A" {
		if w.CheckForMAS(row, col) {
			count++
		}
	}
	return count
}

func (w WordSearch) FindAllXMAS() (wordCode int, xmasCount int) {
	r, c := w.GetBounds()
	wc := 0
	xm := 0
	for i := 0; i <= r; i++ {
		for j := 0; j <= c; j++ {
			wc = wc + w.CheckEachDirection(i, j)
			xm = xm + w.CheckForTwoDiagonalXMAS(i, j)
		}
	}
	return wc, xm
}

func main() {
	w := WordSearchInit("day_04/input.txt")
	wc, xm := w.FindAllXMAS()
	fmt.Printf("Word Count: %d\n", wc)
	fmt.Printf("X-MAS Count: %d\n", xm)
}
