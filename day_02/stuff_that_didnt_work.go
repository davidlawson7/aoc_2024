package day_02

import "fmt"

// The levels are either all increasing or all decreasing.
// Any two adjacent levels differ by at least one and at most three.
func numberOfSafeReports(reports [][]int, withDampener bool) []bool {
	safeReports := make([]bool, 0, len(reports))

	for reportNumber, report := range reports {
		safe, index := determineIfReportIsSafe(report)

		if withDampener && !safe {
			// Second attempt, removes the first bad level and tries again
			fmt.Printf("%d: Removing: %d\n", reportNumber+1, report[index])
			altReport := append(report[:index], report[index+1:]...)
			safe, _ = determineIfReportIsSafe(altReport)
		}
		// Save the reports value. If it didnt change direction, safe will be true by default
		safeReports = append(safeReports, safe)
	}

	return safeReports
}

func determineIfReportIsSafe(report []int) (bool, int) {
	var prev int
	var direction string

	for index, level := range report {
		// First run, skip but save its value
		if index == 0 {
			// Reset everything for a fresh run
			direction = ""
			prev = level
			continue
		}

		// By the second run we know if its increasing or decreasing but not enough to mark as unsafe
		if index == 1 {
			if level < prev {
				if (prev - level) > 3 {
					return false, index // Diff was greater than 3
				}
				direction = "decreasing"
			} else if level > prev {
				if (level - prev) > 3 {
					return false, index // Diff was greater than 3
				}
				direction = "increasing"
			} else {
				// adjacent levels cannot be the same
				return false, index
			}

			prev = level // save the prev value
			continue
		}

		if prev == level {
			return false, index
		}

		// Anything beyond the third point could mark it as unsafe as the direction could change
		if (direction == "increasing" && level < prev) || (direction == "decreasing" && level > prev) {
			// the direction change on us from its initial direction. End run
			return false, index - 1
		}

		if (direction == "increasing" && (level-prev) > 3) || (direction == "decreasing" && (prev-level) > 3) {
			return false, index
		}

		// current level didnt change direction, reset prev to current for next run
		prev = level
	}
	return true, -1
}

// The levels are either all increasing or all decreasing.
// Any two adjacent levels differ by at least one and at most three.
// Note
// 627 is TOO HIGHT
func numberOfSafeReportsDampened(reports [][]int) []bool {
	safeReports := make([]bool, 0, len(reports))

	for _, report := range reports {
		safe := true
		var direction string
		var prev int
		problemDampenerUsed := false

		for index, level := range report {
			// First run, skip but save its value
			if index == 0 {
				// Reset everything for a fresh run
				safe = true
				direction = ""
				prev = level
				continue
			}

			// By the second or third run we know if its increasing or decreasing but not enough to mark as unsafe
			if index == 1 || (index == 2 && direction == "") {
				if level < prev {
					if (prev - level) > 3 {
						// Our 1 chance catch incase a issue arises. Prev remains what was currently set
						// So its like the bad result never happened
						if !problemDampenerUsed {
							problemDampenerUsed = true
							continue // Can recover, continue
						} else {
							safe = false
							break // Diff was greater than 3, cant recover, exit
						}
					}
					direction = "decreasing"
				} else if level > prev {
					if (level - prev) > 3 {
						// Our 1 chance catch incase a issue arises. Prev remains what was currently set
						// So its like the bad result never happened
						if !problemDampenerUsed {
							problemDampenerUsed = true
							continue // Can recover, continue
						} else {
							safe = false
							break // Diff was greater than 3, cant recover, exit
						}
					}
					direction = "increasing"
				} else {
					// Our 1 chance catch incase a issue arises. Prev remains what was currently set
					// So its like the bad result never happened
					if !problemDampenerUsed {
						problemDampenerUsed = true
						continue // Can recover, continue
					} else {
						safe = false
						break // adjacent levels cannot be the same, cant recover, exit
					}
				}

				prev = level // save the prev value
				continue
			}

			// Anything beyond the third point could mark it as unsafe as the direction could change
			if (prev == level) || (direction == "increasing" && level < prev) || (direction == "increasing" && (level-prev) > 3) || (direction == "decreasing" && level > prev) || (direction == "decreasing" && (prev-level) > 3) {
				// Our 1 chance catch incase a issue arises. Prev remains what was currently set
				// So its like the bad result never happened
				if !problemDampenerUsed {
					problemDampenerUsed = true
					continue // Can recover, continue
				} else {
					safe = false
					break // the direction change on us from its initial direction, cant recover, end run
				}
			}

			// current level didnt change direction, reset prev to current for next run
			prev = level
		}

		// Save the reports value. If it didnt change direction, safe will be true by default
		safeReports = append(safeReports, safe)
	}

	return safeReports
}
