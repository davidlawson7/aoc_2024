package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func sum(m []int) int {
	s := 0
	for _, v := range m {
		s += v
	}
	return s
}

func justDoIt(filePath *string) int {
	file, err := os.ReadFile(*filePath) // just pass the file name
	if err != nil {
		log.Fatal(err)
	}
	str := string(file) // convert content to a 'string'

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	matches := re.FindAllString(str, -1)

	values := make([]int, 0, len(matches))

	for _, mul := range matches {
		rmPrefix := strings.Replace(mul, "mul(", "", 1)
		rmPostfix := strings.Replace(rmPrefix, ")", "", 1)
		inputs := strings.Split(rmPostfix, ",")

		if leftNumber, err := strconv.Atoi(inputs[0]); err == nil {
			if rightNumber, err := strconv.Atoi(inputs[1]); err == nil {

				if (leftNumber < 0 || leftNumber > 999) && (rightNumber < 0 || rightNumber > 999) {
					continue
				}
				values = append(values, leftNumber*rightNumber)
			}
		}
	}

	return sum(values)
}

func main() {
	testFile := "day_03/input_test.txt"
	realFile := "day_03/input.txt"

	fmt.Printf("%d\n", justDoIt(&testFile))
	fmt.Printf("%d\n", justDoIt(&realFile))
}
