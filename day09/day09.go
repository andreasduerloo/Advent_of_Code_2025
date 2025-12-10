package day09

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
	inputFile:     "09",
	testInputFile: "09t",
	isTest:        false,
}

func (d day) Solve() {
	file := d.inputFile
	if d.isTest {
		file = d.testInputFile
	}

	coords, err := helpers.SlicesOfInts(file)
	if err != nil {
		fmt.Println("issue getting the input")
	}

	points := convert(coords)

	first := largestArea(points)

	fmt.Println(first)
}
