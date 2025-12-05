package day03

import (
	"strconv"
)

type bank []int

func makeBank(s string) bank {
	ints := make([]int, 0)

	for _, r := range s {
		i, _ := strconv.Atoi(string(r))
		ints = append(ints, i)
	}

	return ints
}

func highestJoltage(b bank, current, needed int) int {
	var highest int

	if needed == 1 {
		for _, batt := range b {
			if 10*current+batt > highest {
				highest = 10*current + batt
			}
		}
	} else {
		var localMax int
		for i, batt := range b[:len(b)-(needed-1)] { // We need enough space at the end for the remaining digits
			if batt > localMax {
				localMax = batt
				option := highestJoltage(b[i+1:], 10*current+batt, needed-1)
				if option > highest {
					highest = option
				}
			}
		}
	}

	return highest
}

func highestJoltageFullRange(b bank, current, needed int) int {
	var highest int

	if needed == 1 {
		for _, batt := range b {
			if 10*current+batt > highest {
				highest = 10*current + batt
			}
		}
	} else {
		var localMax int
		for i, batt := range b { // Barely a performance difference
			if batt > localMax {
				localMax = batt
				option := highestJoltage(b[i+1:], 10*current+batt, needed-1)
				if option > highest {
					highest = option
				}
			}
		}
	}

	return highest
}
