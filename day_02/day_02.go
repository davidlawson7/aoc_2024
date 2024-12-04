package day_02

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Extracts data from the path
func extractData(filePath *string) [][]int {
	file, err := os.Open(*filePath)

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var reports []string

	for fileScanner.Scan() {
		reports = append(reports, fileScanner.Text())
	}

	file.Close()

	r := make([][]int, 0, len(reports))

	for _, report := range reports {
		levels := strings.Fields(report)
		l := make([]int, 0, len(levels))

		for _, levelString := range levels {
			if level, err := strconv.Atoi(levelString); err == nil {
				l = append(l, level)
			}
		}
		r = append(r, l)
	}

	return r
}

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func sumSuccessfulReports(reports []bool) int {
	sum := 0
	for _, isSafe := range reports {
		if isSafe {
			sum += 1
		}
	}
	return sum
}

func IsSorted(report []int) bool {
	return slices.IsSorted(report) || slices.IsSortedFunc(report, func(a, b int) int {
		if a > b {
			return -1
		}

		if b > a {
			return 1
		}

		return 0
	})
}

func IsSafe(report []int) bool {

	if !IsSorted(report) {
		return false
	}

	slices.Sort(report)
	for i := 0; i < len(report); i++ {
		if len(report) == i+1 {
			break // end of report
		}
		// fmt.Println("Compare: %d to %d\n", )

		diff := report[i+1] - report[i]

		if diff == 0 || diff > 3 {
			return false
		}
	}

	return true
}

func IsSafeDampened(report []int) bool {

	if IsSafe(report) {
		return true
	}

	safeWhenDampened := false
	// Not safe, remove values in intervals until either a safe run is found,
	// else its not safe event with dampener

	for index := range report {
		altReport := RemoveIndex(report, index)

		if IsSafe(altReport) {
			return true
		}
	}
	return safeWhenDampened
}

func ProblemOne(reports [][]int) int {
	safeReports := make([]bool, 0, len(reports))

	for _, report := range reports {
		safeReports = append(safeReports, IsSafe(report))
	}

	return sumSuccessfulReports(safeReports)
}

func ProblemTwo(reports [][]int) int {
	safeReports := make([]bool, 0, len(reports))

	for _, report := range reports {
		safeReports = append(safeReports, IsSafeDampened(report))
	}

	return sumSuccessfulReports(safeReports)
}
