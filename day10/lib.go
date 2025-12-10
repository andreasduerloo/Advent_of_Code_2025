package day10

import (
	"advent/helpers"
	"math"
	"regexp"
)

type machine struct {
	goal        int // Let's do some bitwise operations
	joltGoal    []int
	buttons     []button
	joltButtons []joltButton
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
	joltageRe := regexp.MustCompile(`\{.*\}`)

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

		joltButtons := make([]joltButton, 0, len(buttonmatches))

		for _, jb := range buttonmatches {
			joltButtons = append(joltButtons, joltButton{
				effect: helpers.ReGetInts(jb),
			})
		}

		joltageMatches := joltageRe.FindString(line)
		joltageGoal := helpers.ReGetInts(joltageMatches)

		out = append(out, machine{
			goal:        goalInt,
			buttons:     buttons,
			joltGoal:    joltageGoal,
			joltButtons: joltButtons,
		})
	}

	return out
}

func (m machine) findLeastPresses(lights int) int {
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

type joltButton struct {
	effect []int
}

func (jb joltButton) press(values []int) {
	for _, eff := range jb.effect {
		values[eff] += 1
	}
}

type joltState struct {
	presses int
	values  []int
}

func stillPossible(goal, values []int) bool {
	out := true

	for i, g := range goal {
		out = out && (g >= values[i])
	}

	return out
}

func compareValues(goal, values []int) bool {
	out := true

	for i, g := range goal {
		out = out && (g == values[i])
	}

	return out
}

func (m machine) findLeastJoltPresses() int {
	queue := make([]joltState, 0)

	// Add all the buttons to the queue (unless pusing that button once solved it)
	for _, b := range m.joltButtons {
		localValues := make([]int, len(m.joltGoal))

		b.press(localValues)

		if compareValues(m.joltGoal, localValues) {
			return 1
		} else {
			queue = append(queue, joltState{
				presses: 1,
				values:  localValues,
			})
		}
	}

	// Take the top element off the queue, add all button presses to the back of the queue (unless it's fixed)
	current := queue[0]
	queue = queue[1:]

	for {
		for _, b := range m.joltButtons {
			localValues := make([]int, len(current.values))
			copy(localValues, current.values)

			b.press(localValues)

			if compareValues(m.joltGoal, localValues) {
				return current.presses + 1
			}

			// No dice, add to the queue unless one of the values is too high
			if stillPossible(m.joltGoal, current.values) {
				newValues := make([]int, len(localValues))
				copy(newValues, localValues)

				queue = append(queue, joltState{
					presses: current.presses + 1,
					values:  newValues,
				})
			}
		}
		current = queue[0]
		queue = queue[1:]
	}
}
