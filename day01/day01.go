package day01

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
	inputFile:     "01",
	testInputFile: "01t",
	isTest:        false,
}

type lock struct {
	position   int
	atZero     int
	passedZero int
}

func (l *lock) turn(direction rune, number int) {
	var newPos int

	switch direction {
	case 'R':
		newPos = l.position + number

		l.passedZero += newPos / 100 // Full rotations
		newPos = newPos % 100

		if newPos == 0 {
			l.atZero++
		}

	case 'L':
		for number > 100 { // Full rotations
			number -= 100
			l.passedZero++
		}

		newPos = l.position - number

		if newPos == 0 || newPos == 100 || newPos == -100 {
			l.passedZero++
			l.atZero++
			newPos = 0
		}

		if newPos < 0 && newPos != -100 && l.position != 0 {
			l.passedZero++
			newPos += 100
		}

		if newPos < 0 && newPos != -100 && l.position == 0 {
			newPos += 100
		}
	}

	l.position = newPos
}

func (d day) Solve() {
	file := d.inputFile
	if d.isTest {
		file = d.testInputFile
	}

	fullStrings, err := helpers.SliceOfStrings(file)
	if err != nil {
		fmt.Println("issue getting string input")
	}

	ints, err := helpers.SliceOfInts(file)
	if err != nil {
		fmt.Println("issue getting int input")
	}

	safe := lock{
		position: 50,
	}

	for i, dir := range fullStrings {
		safe.turn(rune(dir[0]), ints[i])
	}

	fmt.Println(safe.atZero, safe.passedZero)
}
