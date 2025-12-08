package day08

import (
	"advent/helpers"
	"fmt"
)

type day struct {
	inputFile     string
	testInputFile string
	isTest        bool
	connections   int
}

var Config = day{
	inputFile:     "08",
	testInputFile: "08t",
	isTest:        false,
	connections:   1000,
}

func (d day) Solve() {
	file := d.inputFile
	if d.isTest {
		file = d.testInputFile
	}

	coords, err := helpers.SlicesOfInts(file)
	if err != nil {
		fmt.Println("issue getting the input")
	}

	// Just shamelessly front load all distance calculations and sort by that?
	points := transfromToPoints(coords)
	edges := getEdges(points)

	// Feels fast enough to me!

	pl := newPlayground()

	connections := 0
	var lastEdge edge

	for !pl.allConnected(points) || connections < d.connections {
		pl.union(edges[connections])
		lastEdge = edges[connections]

		if connections == d.connections-1 {
			fmt.Println(threeLargest(pl.circuits))
		}

		connections++
	}

	fmt.Println(lastEdge.p1.x * lastEdge.p2.x)
}
