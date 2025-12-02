package day02

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
	inputFile:     "02",
	testInputFile: "02t",
	isTest:        false,
}

func (d day) Solve() {
	file := d.inputFile
	if d.isTest {
		file = d.testInputFile
	}

	ints, err := helpers.SliceOfPureInts(file)
	if err != nil {
		fmt.Println("issue getting string input")
	}

	// First star
	var first int

	for i := 0; i < len(ints)-1; i += 2 {
		first += scanRange(ints[i], ints[i+1])
	}

	// Second star
	var second int

	for i := 0; i < len(ints)-1; i += 2 {
		second += secondScanRange(ints[i], ints[i+1])
	}

	fmt.Println(first, second)
}
