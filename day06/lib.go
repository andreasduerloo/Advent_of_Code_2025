package day06

import (
	"advent/helpers"
	"regexp"
	"strconv"
	"strings"
)

// Adapt the helper function to include + and *
func getOperators(day string) ([]string, error) {
	s, err := helpers.GetInputFromString(day)
	out := []string{}

	if err != nil {
		return out, err
	}

	lines := strings.Split(strings.TrimSpace(s), "\n")

	for _, line := range lines {
		out = append(out, reGetOps(line)...)
	}

	return out, nil
}

func reGetOps(s string) []string {
	re := regexp.MustCompile(`[+*]`)
	matches := re.FindAllString(s, -1)

	return matches
}

func calculate(data [][]int, operators []string) int {
	var out int

	for col := 0; col < len(operators); col++ {
		var localResult int
		for row, _ := range data {
			if row == 0 {
				localResult = data[row][col]
			} else {
				switch operators[col] {
				case "+":
					localResult += data[row][col]
				case "*":
					localResult *= data[row][col]
				}
			}
		}

		out += localResult
	}

	return out
}

// Second star

type operation struct {
	operands  []int
	operation string
}

func newOperation() *operation {
	return &operation{
		operands:  []int{},
		operation: "",
	}
}

func (o *operation) calculate() int {
	var out int
	for i, num := range o.operands {
		if i == 0 {
			out = num
		} else {
			switch o.operation {
			case "+":
				out += num
			case "*":
				out *= num
			}
		}
	}

	return out
}

func transform(data []string) []*operation {
	out := make([]*operation, 0)

	current := newOperation()

	rows := len(data)

	// Scan through all strings at the same time, until we hit a line that is only spaces
	for i := 0; i < len(data[0]); i++ {
		allSpaces := true
		number := 0

		for r := 0; r < rows; r++ {
			if string(data[r][i]) != " " {
				allSpaces = false
			}
		}

		if allSpaces {
			out = append(out, current)
			current = newOperation()
			continue
		}

		for r := 0; r < rows; r++ {
			switch string(data[r][i]) {
			case "+":
				current.operation = "+"
			case "*":
				current.operation = "*"
			case " ":
				continue
			default:
				if number == 0 {
					number, _ = strconv.Atoi(string(data[r][i]))
				} else {
					newNumber, _ := strconv.Atoi(string(data[r][i]))
					number = 10*number + newNumber
				}
			}
		}
		current.operands = append(current.operands, number)
	}
	out = append(out, current)
	return out
}
