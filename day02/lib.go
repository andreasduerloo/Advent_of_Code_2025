package day02

import (
	"math"
	"strconv"
	"strings"
)

// First star

func inRange(value, low, high int) bool {
	return value >= low && value <= high
}

func lowestCandidate(value int) int {
	stringWise := strconv.Itoa(value)

	// The input number has an even length -> the lowest candidate is the first "half" (stringwise) of that number
	if len(stringWise)%2 == 0 {
		firstHalf := stringWise[0 : len(stringWise)/2]

		out, _ := strconv.Atoi(firstHalf)

		return out
	}

	// The input number has an odd length -> the lowest candidate is the first half of next power of ten
	NumberOfDigits := len(stringWise)
	nextEven := int(math.Pow(10, float64(NumberOfDigits)))

	return lowestCandidate(nextEven)
}

func doubleCandidate(value int) int {
	stringWise := strconv.Itoa(value)
	newString := stringWise + stringWise

	out, _ := strconv.Atoi(newString) // Should logically never error

	return out
}

func scanRange(low, max int) int {
	var out int

	candidate := lowestCandidate(low)

	for doubleCandidate(candidate) <= max {
		if inRange(doubleCandidate(candidate), low, max) {
			out += doubleCandidate(candidate)
		}
		candidate++
	}

	return out
}

// Second star

type set map[int]struct{}

/*
func (s set) contains(i int) bool {
	_, present := s[i]
	return present
}
*/

func (s set) add(i int) {
	s[i] = struct{}{}
}

func (s set) sum() int {
	var out int
	for k, _ := range s {
		out += k
	}
	return out
}

func secondScanRange(low, max int) int {
	ids := set{} // Use a set so we don't double-count

	// Re-use what already works: the "classic" invalid ID's
	candidate := lowestCandidate(low)

	for doubleCandidate(candidate) <= max {
		if inRange(doubleCandidate(candidate), low, max) {
			ids.add(doubleCandidate(candidate))
		}
		candidate++
	}

	// The longest input numbers are 10 digits, so we need to check lengths (periods) of 1, 2, 3, 4, and 5
	// 4 and 5 we can entirely forget, because they will always be 'classic' invalid ID's and will already have been caught
	// This means we need to check for repetition length 1, 2, and 3

	// Repetition length of 1 (i.e. 111)
	rep1(ids, low, max)

	// Repetition length of 2 (i.e. 121212)
	rep2(ids, low, max)

	// Repetition length of 3 (i.e. 123123123)
	rep3(ids, low, max)

	return ids.sum()
}

func rep1(s set, min, max int) {
	minStr := strconv.Itoa(min)
	maxStr := strconv.Itoa(max)

	candidateLen := len(minStr) // bugs on 2 - 16
	if candidateLen == 1 {
		candidateLen++
	}

	for candidateLen <= len(maxStr) {
		baseCandidate, _ := strconv.Atoi(strings.Repeat("1", candidateLen))
		for i := 1; i <= 9; i++ { // As long as candidate <= max - TODO
			candidate := i * baseCandidate
			if inRange(candidate, min, max) {
				s.add(candidate)
			}
		}
		candidateLen++
	}
}

func rep2(s set, min, max int) {
	// This is only relevant for even-length numbers
	minStr := strconv.Itoa(min)
	maxStr := strconv.Itoa(max)

	for candidateLen := len(minStr); candidateLen <= len(maxStr); candidateLen++ {
		if candidateLen < 4 || candidateLen%2 != 0 {
			continue // Not the right length, ignore
		}

		for i := 10; i <= 99; i++ { // Brute force, unfortunately, needs to be pruned aggressively
			candidate, _ := strconv.Atoi(strings.Repeat(strconv.Itoa(i), candidateLen/2))
			if inRange(candidate, min, max) {
				s.add(candidate)
			}
		}
	}
}

func rep3(s set, min, max int) {
	// This is only relevant for numbers of length 6 or 9
	minStr := strconv.Itoa(min)
	maxStr := strconv.Itoa(max)

	for candidateLen := len(minStr); candidateLen <= len(maxStr); candidateLen++ {
		if candidateLen == 6 || candidateLen == 9 {
			for i := 100; i <= 999; i++ { // Brute force, unfortunately, needs to be pruned aggressively
				candidate, _ := strconv.Atoi(strings.Repeat(strconv.Itoa(i), candidateLen/3))
				if inRange(candidate, min, max) {
					s.add(candidate)
				}
			}
		}
	}
}
