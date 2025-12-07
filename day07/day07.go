package day07

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
	inputFile:     "07",
	testInputFile: "07t",
	isTest:        false,
}

func (d day) Solve() {
	file := d.inputFile
	if d.isTest {
		file = d.testInputFile
	}

	data, err := helpers.SliceOfStrings(file)
	if err != nil {
		fmt.Println("issue getting the input")
	}

	first := countSplits(data)
	second := totalPaths(data)

	fmt.Println(first, second)
}
