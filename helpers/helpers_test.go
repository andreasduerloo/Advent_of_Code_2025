package helpers

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Chdir("..")
	exitval := m.Run()
	os.Exit(exitval)
}

func Test_GetInputFromString(t *testing.T) {
	result, err := GetInputFromString("00")

	if err != nil {
		t.Error("error reading file:", err)
	}

	if len(result) != 27 {
		t.Error("file length is not correct")
	}
}

func Test_SliceOfInts(t *testing.T) {
	result, err := SliceOfInts("00")

	if err != nil {
		t.Error(err)
	}

	if result[0] != 1 || result[len(result)-1] != -5 || len(result) != 5 {
		t.Error("wrong ints")
	}
}

func Test_SlicesOfInts(t *testing.T) {
	result, err := SlicesOfInts("00")

	if err != nil {
		t.Error(err)
	}

	if len(result[0]) != 1 || result[0][0] != 1 || result[len(result)-1][len(result[len(result)-1])-1] != -5 {
		t.Error("wrong ints")
	}
}

func Test_SliceOfStrings(t *testing.T) {
	result, err := SliceOfStrings("00")

	if err != nil {
		t.Error(err)
	}

	if len(result) != 3 || result[0] != "a, b, 1" || result[len(result)-1] != "xyz, 15, -5" {
		t.Error("wrong strings")
	}
}

/*
func Test_SlicesOfStrings(t *testing.T) {
	result, err := SlicesOfStrings("00")

	if err != nil {
		t.Error(err)
	}
}
*/
