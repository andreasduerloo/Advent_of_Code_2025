package day10

import (
	"advent/helpers"
	"math"
	"regexp"
)

type machine struct {
	goal    int // Let's do some bitwise operations
	buttons []button
}

type button struct {
	effect int
}

func (b button) press(lights int) int {
	return b.effect ^ lights
}

type state struct {
	presses       int
	currentLights int
}

func transform(lines []string) []machine {
	lightsRe := regexp.MustCompile(`[#.]+`)
	buttonRe := regexp.MustCompile(`\([^(]*\)`)

	out := make([]machine, 0, len(lines))

	for _, line := range lines {
		goal := lightsRe.FindString(line)
		var goalInt int

		for i, r := range goal {
			switch r {
			case '.':
				continue
			case '#':
				goalInt += int(math.Pow(2, float64(i))) // We flip the lights this way
			}
		}

		buttonmatches := buttonRe.FindAllString(line, -1)

		buttons := make([]button, 0, len(buttonmatches))

		for _, b := range buttonmatches {
			numbers := helpers.ReGetInts(b)
			var effect int

			for _, n := range numbers {
				effect += int(math.Pow(2, float64(n))) // We flip the lights this way
			}

			buttons = append(buttons, button{effect: effect})
		}

		out = append(out, machine{
			goal:    goalInt,
			buttons: buttons,
		})
	}

	return out
}

func (m machine) findLeastPresses(lights, presses int) int {
	// Work with a queue, with recursion it will go way too deep on the first button first
	queue := make([]state, 0)

	// Add all the buttons to the queue (unless pusing that button once solved it)
	for _, b := range m.buttons {
		if b.press(lights) == m.goal {
			return 1
		} else {
			queue = append(queue, state{
				presses:       1,
				currentLights: b.press(lights),
			})
		}
	}

	// Take the top element off the queue, add all button presses to the back of the queue (unless it's fixed)
	current := queue[0]
	queue = queue[1:]

	for {
		for _, b := range m.buttons {
			if b.press(current.currentLights) == m.goal {
				return current.presses + 1
			}

			// No dice, add to the queue
			queue = append(queue, state{
				presses:       current.presses + 1,
				currentLights: b.press(current.currentLights),
			})
		}
		current = queue[0]
		queue = queue[1:]
	}
}
