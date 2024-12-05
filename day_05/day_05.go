package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type PrintQueue struct {
	order   map[string][]string
	updates [][]string
}

func atoi(s string) int {
	value, _ := strconv.Atoi(s)
	return value
}

func (p *PrintQueue) ResortIncorrectUpdate(update []string) []string {
	slices.SortFunc(update, func(a, b string) int {
		if slices.Contains(p.order[a], b) {
			return 1
		}

		if slices.Contains(p.order[b], a) {
			return -1
		}

		return 0
	})

	return update
}

func InitPrintQueue(filePath string) *PrintQueue {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	strNormalized := regexp.MustCompile("\r\n").ReplaceAllString(string(file), "\n")
	fields := regexp.MustCompile(`\n\s*\n`).Split(strNormalized, -1)

	order := make(map[string][]string)
	for _, rule := range strings.Fields(fields[0]) {
		xy := strings.Split(rule, "|")
		_, exists := order[xy[0]]

		if !exists {
			// field doesnt exist yet, create it
			var yValues []string
			order[xy[0]] = append(yValues, xy[1])
		} else {
			// Does exist, append the most recent value if the slice doesnt already have it
			if !slices.Contains(order[xy[0]], xy[1]) {
				order[xy[0]] = append(order[xy[0]], xy[1])
			}
		}
	}

	var updates [][]string
	for _, update := range strings.Fields(fields[1]) {
		updates = append(updates, strings.Split(update, ","))
	}

	return &PrintQueue{
		order:   order,
		updates: updates,
	}
}

func (p *PrintQueue) checkUpdate(update []string) bool {
	for i, page := range update {
		for j := 0; j < i; j++ {
			pagesThatMustComeBefore := p.order[page]
			if slices.Contains(pagesThatMustComeBefore, update[j]) {
				return false
			}
		}
	}

	return true
}

func (p *PrintQueue) GetCorrectUpdates() (int, int) {
	sum := 0
	sumBad := 0
	for _, update := range p.updates {
		if p.checkUpdate(update) {
			l := len(update)

			if l%2 == 0 {
				sum += (atoi(update[l/2-1]) + atoi(update[l/2])) / 2
			} else {
				sum += atoi(update[l/2])
			}
		} else {
			resorted := p.ResortIncorrectUpdate(update)

			l := len(resorted)

			if l%2 == 0 {
				sumBad += (atoi(resorted[l/2-1]) + atoi(resorted[l/2])) / 2
			} else {
				sumBad += atoi(resorted[l/2])
			}
		}
	}

	return sum, sumBad
}

func main() {
	pq := InitPrintQueue("day_05/input.txt")
	p1, p2 := pq.GetCorrectUpdates()

	fmt.Printf("Part 1: %d\n", p1)
	fmt.Printf("Part 2: %d\n", p2)
}
