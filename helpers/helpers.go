package helpers

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Solver interface {
	Solve()
}

// Functions to return slices or slices of slices (blocks)
func GetInputFromString(day string) (string, error) {
	path := "./inputs/" + day + ".txt"

	input, err := os.ReadFile(path)
	if err != nil {
		return "", err
	} else {
		return string(input), nil
	}
}

func ReGetInts(s string) []int {
	re := regexp.MustCompile(`-?[0-9]+`)
	matches := re.FindAllString(s, -1)

	ints := make([]int, 0)

	for _, match := range matches {
		val, err := strconv.Atoi(match)
		if err != nil {
			continue
		}
		ints = append(ints, val)
	}

	return ints
}

func SliceOfInts(day string) ([]int, error) {
	s, err := GetInputFromString(day)
	out := []int{}

	if err != nil {
		return out, err
	}

	out = ReGetInts(s)

	return out, nil
}

func SlicesOfInts(day string) ([][]int, error) {
	s, err := GetInputFromString(day)
	out := [][]int{}

	if err != nil {
		return out, err
	}

	lines := strings.Split(strings.TrimSpace(s), "\n")

	for _, line := range lines {
		out = append(out, ReGetInts(line))
	}

	return out, nil
}

func SliceOfStrings(day string) ([]string, error) {
	s, err := GetInputFromString(day)
	out := []string{}

	if err != nil {
		return out, err
	}

	out = strings.Split(strings.TrimSpace(s), "\n")

	return out, nil
}

func SlicesOfStrings(day string) ([][]string, error) {
	s, err := GetInputFromString(day)
	out := [][]string{}

	if err != nil {
		return out, err
	}

	blocks := strings.Split(strings.TrimSpace(s), "\n\n")

	for _, block := range blocks {
		out = append(out, strings.Split((block), "\n"))
	}

	return out, nil
}
