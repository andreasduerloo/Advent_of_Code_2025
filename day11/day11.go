package day11

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
	inputFile:     "11",
	testInputFile: "11t",
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

	serverRoom := parseInput(lines)

	first := serverRoom.countPaths("you", "out")

	fmt.Println(first)
}
