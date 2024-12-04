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

func GetInputs(mul string) (int, int) {
	rmPrefix := strings.Replace(mul, "mul(", "", 1)
	rmPostfix := strings.Replace(rmPrefix, ")", "", 1)
	inputs := strings.Split(rmPostfix, ",")

	if leftNumber, err := strconv.Atoi(inputs[0]); err == nil {
		if rightNumber, err := strconv.Atoi(inputs[1]); err == nil {
			if (leftNumber < 0 || leftNumber > 999) && (rightNumber < 0 || rightNumber > 999) {
				return 0, 0
			}
			return leftNumber, rightNumber
		}
	}
	return 0, 0
}

func JustDoIt(filePath *string) int {
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

func JustDoItTwo(filePath *string) int {
	file, err := os.ReadFile(*filePath) // just pass the file name
	if err != nil {
		log.Fatal(err)
	}
	str := string(file) // convert content to a 'string'
	re := regexp.MustCompile(`(mul\((\d+),(\d+)\)|don't\(\)|do\(\))`)

	enabled := true
	matches := re.FindAllString(str, -1)
	values := make([]int, 0, len(matches))

	for _, command := range matches {

		if command == "do()" {
			enabled = true
			continue
		} else if command == "don't()" {
			enabled = false
			continue
		}

		if enabled {
			left, right := GetInputs(command)
			values = append(values, left*right)
		}
	}
	return sum(values)
}

func ProblemOne() {
	testFile := "day_03/input_test.txt"
	realFile := "day_03/input.txt"

	fmt.Printf("%d\n", JustDoIt(&testFile))
	fmt.Printf("%d\n", JustDoIt(&realFile))
}

func ProblemTwo() {
	testFile := "day_03/input_test_2.txt"
	realFile := "day_03/input.txt"

	fmt.Printf("%d\n", JustDoItTwo(&testFile))
	fmt.Printf("%d\n", JustDoItTwo(&realFile))

}

func main() {
	ProblemTwo()
}
