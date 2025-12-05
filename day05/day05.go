package day05

import (
	"advent/helpers"
	"fmt"
)

type day struct {
	inputFile     string
	testInputFile string
	isTest        bool
}

var Config = day{
	inputFile:     "05",
	testInputFile: "05t",
	isTest:        false,
}

func (d day) Solve() {
	file := d.inputFile
	if d.isTest {
		file = d.testInputFile
	}

	data, err := helpers.SlicesOfInts(file)
	if err != nil {
		fmt.Println("issue transforming the input")
	}

	// First star: naive way, without unions
	ranges := make([]idRange, 0)
	ids := make([]int, 0)

	for _, elem := range data {
		if len(elem) == 2 {
			ranges = append(ranges, makeRange(elem))
		}
		if len(elem) == 1 {
			ids = append(ids, elem[0])
		}
	}

	var first int
	for _, item := range ids {
		for _, r := range ranges {
			if r.contains(item) {
				// fmt.Println("Item", item, "is in range", r)
				first++
				break
			}
		}
	}

	// Second star, as many unions as we can, then add the lengths
	ranges = compress(ranges)

	lengths := helpers.MapSlice(ranges, idRange.length)
	second := helpers.ReduceSlice(lengths, func(l, acc int) int { return l + acc })

	fmt.Println(first, second)
}
