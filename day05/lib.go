package day05

import "slices"

type idRange struct {
	low  int
	high int
}

func (r idRange) contains(item int) bool {
	return item >= r.low && item <= r.high
}

func makeRange(ints []int) idRange {
	return idRange{
		low:  ints[0],
		high: -ints[1],
	}
}

func overlap(r1, r2 idRange) bool {
	return !(r1.high < r2.low || r1.low > r2.high)
}

func union(r1, r2 idRange) idRange {
	var newLow, newHigh int

	if r1.low <= r2.low {
		newLow = r1.low
	} else {
		newLow = r2.low
	}

	if r1.high >= r2.high {
		newHigh = r1.high
	} else {
		newHigh = r2.high
	}

	return idRange{
		low:  newLow,
		high: newHigh,
	}
}

func (r idRange) length() int {
	return r.high - r.low + 1 // inclusive
}

func compress(idranges []idRange) []idRange {
	slices.SortFunc(idranges, func(r1, r2 idRange) int { // Sort by the lower bound
		return r1.low - r2.low
	})

	out := make([]idRange, 0) // Merge in one pass
	current := idranges[0]

	for i := 1; i < len(idranges); i++ {
		neighbor := idranges[i]

		if overlap(current, neighbor) {
			current = union(current, neighbor) // Current grows, but is not appended
		} else {
			out = append(out, current) // Current is appended
			current = neighbor
		}
	}

	out = append(out, current)
	return out
}
