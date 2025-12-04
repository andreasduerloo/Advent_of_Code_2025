package day03

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
	inputFile:     "03",
	testInputFile: "03t",
	isTest:        false,
}

func (d day) Solve() {
	file := d.inputFile
	if d.isTest {
		file = d.testInputFile
	}

	banks, err := helpers.SliceOfStrings(file)
	if err != nil {
		fmt.Println("issue getting string input")
	}

	// First star
	var first int

	for _, b := range banks {
		first += highestJoltage(makeBank(b), 0, 2)
	}

	fmt.Println(first)

	// Second star
	var second int

	for _, b := range banks {
		second += highestJoltage(makeBank(b), 0, 12) // ugly as sin, but works thanks to the localMax (pruning the regression tree)
	}

	fmt.Println(second)
}
