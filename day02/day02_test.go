package day02

import (
	"testing"
)

func TestLowestCandidateOdd(t *testing.T) {
	result := lowestCandidate(12345)

	if result != 100 {
		t.Error("wrong lowest candidate for an odd-length number")
	}
}

func TestLowestCandidateEven(t *testing.T) {
	result := lowestCandidate(2949)

	if result != 29 {
		t.Error("wrong lowest candidate for an even-length number")
	}
}

func TestScanRange(t *testing.T) {
	result := scanRange(655, 1102)

	if result != 1010 {
		t.Error("wrong result on range odd - even")
	}

	result = scanRange(2949, 4431)

	if result != 47268 {
		t.Error("wrong result on range even - even. Received:", result)
	}
}

func TestDouble(t *testing.T) {
	result := doubleCandidate(29)

	if result != 2929 {
		t.Error("wrong doubling")
	}
}
