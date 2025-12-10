package day10

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
	inputFile:     "10",
	testInputFile: "10t",
	isTest:        false,
}

func (d day) Solve() {
	file := d.inputFile
	if d.isTest {
		file = d.testInputFile
	}

	lines, err := helpers.SliceOfStrings(file)
	if err != nil {
		fmt.Println("issue getting the input")
	}

	machines := transform(lines)

	var presses int

	for _, m := range machines {
		presses += m.findLeastPresses(0)
	}

	fmt.Println(presses)

	/*
		var first, second int

		for _, m := range machines {
			first += m.findLeastPresses(0)
			second += m.findLeastJoltPresses()
			fmt.Println("Found one")
		}

		fmt.Println(first, second)
	*/
}
