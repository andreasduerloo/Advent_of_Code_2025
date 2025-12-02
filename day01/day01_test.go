package day01

import (
	"testing"
)

func TestSimpleRight(t *testing.T) {
	safe := lock{
		position: 50,
	}

	safe.turn('R', 5)

	if safe.position != 55 || safe.atZero != 0 || safe.passedZero != 0 {
		t.Error("result not correct after turning R5")
	}
}

func TestSimpleLeft(t *testing.T) {
	safe := lock{
		position: 50,
	}

	safe.turn('L', 5)

	if safe.position != 45 || safe.atZero != 0 || safe.passedZero != 0 {
		t.Error("result not correct after turning L5")
	}
}

func TestRightEndAtZero(t *testing.T) {
	safe := lock{
		position: 50,
	}

	safe.turn('R', 50)

	if safe.position != 0 {
		t.Error("position is not 0")
	}
	if safe.atZero != 1 {
		t.Error("atzero is not 1")
	}
	if safe.passedZero != 1 {
		t.Error("passedzero is not 1")
	}
}

func TestLeftEndAtZero(t *testing.T) {
	safe := lock{
		position: 50,
	}

	safe.turn('L', 50)

	if safe.position != 0 {
		t.Error("position is not 0")
	}
	if safe.atZero != 1 {
		t.Error("atzero is not 1")
	}
	if safe.passedZero != 1 {
		t.Error("passedzero is not 1")
	}
}

func TestMultipleRight(t *testing.T) {
	safe := lock{
		position: 50,
	}

	safe.turn('R', 200)

	if safe.position != 50 {
		t.Error("position is not 50")
	}
	if safe.atZero != 0 {
		t.Error("atzero is not 0")
	}
	if safe.passedZero != 2 {
		t.Error("passedzero is not 2")
	}
}

func TestMultipleLeft(t *testing.T) {
	safe := lock{
		position: 50,
	}

	safe.turn('L', 200)

	if safe.position != 50 {
		t.Error("position is not 50")
	}
	if safe.atZero != 0 {
		t.Error("atzero is not 0")
	}
	if safe.passedZero != 2 {
		t.Error("passedzero is not 2")
	}
}

func TestMultipleRightEndAtZero(t *testing.T) {
	safe := lock{
		position: 50,
	}

	safe.turn('R', 250)

	if safe.position != 0 {
		t.Error("position is not 0")
	}
	if safe.atZero != 1 {
		t.Error("atzero is not 1")
	}
	if safe.passedZero != 3 {
		t.Error("passedzero is not 3")
	}
}

func TestMultipleLeftEndAtZero(t *testing.T) {
	safe := lock{
		position: 50,
	}

	safe.turn('L', 250)

	if safe.position != 0 {
		t.Error("position is not 0")
	}
	if safe.atZero != 1 {
		t.Error("atzero is not 1")
	}
	if safe.passedZero != 3 {
		t.Error("passedzero is not 3")
	}
}

func TestStopAtZeroMultipleRight(t *testing.T) {
	safe := lock{
		position: 50,
	}

	safe.turn('R', 50)
	safe.turn('R', 250)

	if safe.position != 50 {
		t.Error("position is not 50")
	}
	if safe.atZero != 1 {
		t.Error("atzero is not 1")
	}
	if safe.passedZero != 3 {
		t.Error("passedzero is not 3")
	}
}

func TestStopAtZeroMultipleRightEndAtZero(t *testing.T) {
	safe := lock{
		position: 50,
	}

	safe.turn('R', 50)
	safe.turn('R', 250)
	safe.turn('R', 50)

	if safe.position != 0 {
		t.Error("position is not 0")
	}
	if safe.atZero != 2 {
		t.Error("atzero is not 2")
	}
	if safe.passedZero != 4 {
		t.Error("passedzero is not 4")
	}
}

func TestStopAtZeroMultipleLeft(t *testing.T) {
	safe := lock{
		position: 50,
	}

	safe.turn('L', 50)
	safe.turn('L', 250)

	if safe.position != 50 {
		t.Error("position is not 50")
	}
	if safe.atZero != 1 {
		t.Error("atzero is not 1")
	}
	if safe.passedZero != 3 {
		t.Error("passedzero is not 3")
	}
}

func TestStopAtZeroMultipleLeftEndAtZero(t *testing.T) {
	safe := lock{
		position: 50,
	}

	safe.turn('L', 50)
	safe.turn('L', 250)
	safe.turn('L', 50)

	if safe.position != 0 {
		t.Error("position is not 0")
	}
	if safe.atZero != 2 {
		t.Error("atzero is not 1")
	}
	if safe.passedZero != 4 {
		t.Error("passedzero is not 4")
	}
}
