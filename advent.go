package main

import (
	"advent/day01"
	"advent/day02"
	"advent/day03"
	"advent/day04"
	"advent/day05"
	"advent/day06"
	"advent/day07"
	"advent/day08"
	"advent/day09"
	"advent/day10"
	"advent/day11"
	"advent/helpers"
	"fmt"
	"os"
	"strconv"
)

var solutions = []helpers.Solver{
	day01.Config,
	day02.Config,
	day03.Config,
	day04.Config,
	day05.Config,
	day06.Config,
	day07.Config,
	day08.Config,
	day09.Config,
	day10.Config,
	day11.Config,
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No argument was passed - exiting.")
		return
	}

	day, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("The argument is not an integer - exiting.")
	}

	if day > 0 && day <= len(solutions) {
		fmt.Println("Solutions for day", day)
		solutions[day-1].Solve()
	} else {
		fmt.Println("That's not a valid day.")
	}
}
