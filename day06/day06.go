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
		fmt.Println("issue getting the input")
	}
	data = data[:len(data)-1] // Trim the last, empty slice

	operators, err := getOperators(file)

	first := calculate(data, operators)

	// For the second star we have to start all over with the input
	dataSecondStar, err := helpers.SliceOfStringsNoTrim(file)
	if err != nil {
		fmt.Println("issue getting the input")
	}

	operations := transform(dataSecondStar)

	results := helpers.MapSlice(operations, (*operation).calculate)
	second := helpers.ReduceSlice(results, func(r, acc int) int { return r + acc })

	fmt.Println(first, second)
}
