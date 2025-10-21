package helpers

import (
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Functions to return slices or slices of slices (blocks)
func GetInput(day int) (string, error) {
	if day > 0 && day <= 25 {
		dayString := strconv.Itoa(day)

		if len(dayString) == 1 {
			dayString = "0" + dayString
		}

		path := "./inputs/" + dayString + ".txt"

		input, err := os.ReadFile(path)
		if err != nil {
			return "", err
		} else {
			return string(input), nil
		}
	} else {
		return "", errors.New("not a valid day")
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

func SliceOfInts(day int) ([]int, error) {
	s, err := GetInput(day)
	out := []int{}

	if err != nil {
		return out, err
	}

	out = ReGetInts(s)

	return out, nil
}

func SlicesOfInts(day int) ([][]int, error) {
	s, err := GetInput(day)
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

func SliceOfStrings(day int) ([]string, error) {
	s, err := GetInput(day)
	out := []string{}

	if err != nil {
		return out, err
	}

	out = strings.Split(strings.TrimSpace(s), "\n")

	return out, nil
}

func SlicesOfStrings(day int) ([][]string, error) {
	s, err := GetInput(day)
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
