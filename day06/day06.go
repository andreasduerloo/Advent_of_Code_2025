package day06

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
	inputFile:     "06",
	testInputFile: "06t",
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

	fmt.Println(data)
}
