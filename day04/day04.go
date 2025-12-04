package day04

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
	inputFile:     "04",
	testInputFile: "04t",
	isTest:        false,
}

func (d day) Solve() {
	file := d.inputFile
	if d.isTest {
		file = d.testInputFile
	}

	grid, err := helpers.SliceOfStrings(file)
	if err != nil {
		fmt.Println("issue getting string input")
	}

	first, _, _ := scanForAccess(grid)
	second := reduce(grid)

	fmt.Println(first, second)
}
