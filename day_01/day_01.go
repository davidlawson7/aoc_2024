package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Extracts data from the path
func extractData(filePath *string) (*[]int, *[]int) {
	file, err := os.Open(*filePath)

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	file.Close()

	left := make([]int, 0, len(fileLines))
	right := make([]int, 0, len(fileLines))

	for _, line := range fileLines {
		values := strings.Fields(line)

		if leftNumber, err := strconv.Atoi(values[0]); err == nil {
			left = append(left, leftNumber)
		}

		if rightNumber, err := strconv.Atoi(values[1]); err == nil {
			right = append(right, rightNumber)
		}
	}

	// slices.Sort(left)
	// slices.Sort(right)

	return &left, &right
}

func calcDistances(left []int, right []int) *[]int {
	if len(left) != len(right) {
		log.Fatal("Left and right side aren't the same length")
	}

	slices.Sort(left)
	slices.Sort(right)

	d := make([]int, 0, len(left))
	for i := 0; i < len(left); i++ {
		if left[i] > right[i] {
			d = append(d, left[i]-right[i])
		} else {
			d = append(d, right[i]-left[i])
		}
	}

	return &d
}

func totalDistance(distances []int) int {
	sum := 0
	for _, distance := range distances {
		sum += distance
	}
	return sum
}

func main() {
	filePath := os.Args[1]
	left, right := extractData(&filePath)
	distances := calcDistances(*left, *right)
	td := totalDistance(*distances)

	fmt.Printf("Total Distance: %d\n", td)
}
