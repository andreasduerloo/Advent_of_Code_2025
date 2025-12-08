package day08

import (
	"math"
	"slices"
	"sort"
)

type point struct {
	x       int
	y       int
	z       int
	circuit int
}

type circuit struct {
	id     int
	points []*point
}

type edge struct {
	p1       *point
	p2       *point
	distance float64
}

func distance(a, b *point) float64 {
	return math.Sqrt(math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2) + math.Pow(float64(a.z-b.z), 2))
}

func newPoint(coords []int) *point {
	return &point{
		x: coords[0],
		y: coords[1],
		z: coords[2],
	}
}

func newEdge(p1, p2 *point) edge {
	return edge{
		p1:       p1,
		p2:       p2,
		distance: distance(p1, p2),
	}
}

func transfromToPoints(coords [][]int) []*point {
	out := make([]*point, 0, len(coords))
	for _, coordset := range coords {
		out = append(out, newPoint(coordset))
	}

	return out
}

func getEdges(points []*point) []edge {
	// Unordered pairs with no self-pairing: n(n-1)/2 combinations
	out := make([]edge, 0, len(points)*(len(points)-1)/2) // Avoid re-allocations, the slice is going to be massive: 499_500 elements

	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			out = append(out, newEdge(points[i], points[j]))
		}
	}

	slices.SortFunc(out, func(e1, e2 edge) int {
		return int(e1.distance - e2.distance)
	})

	return out
}

type playground struct {
	circuits      map[int]circuit
	nextAvailable int
}

func newPlayground() playground {
	return playground{
		circuits:      make(map[int]circuit),
		nextAvailable: 1,
	}
}

func (p *playground) union(e edge) {
	// Both points have no circuit, make a new one
	if e.p1.circuit == 0 && e.p2.circuit == 0 {
		// Make the circuit
		newCircuit := circuit{
			id:     p.nextAvailable,
			points: []*point{e.p1, e.p2},
		}
		p.circuits[p.nextAvailable] = newCircuit

		// Set the circuit on the points
		e.p1.circuit = p.nextAvailable
		e.p2.circuit = p.nextAvailable

		p.nextAvailable++
		return
	}

	// One point has a circuit, join it with the other point
	if (e.p1.circuit == 0) != (e.p2.circuit == 0) { // This is an XOR - one point has a circuit, the other doesn't
		if e.p1.circuit == 0 {
			e.p1.circuit = e.p2.circuit

			currentCircuit := p.circuits[e.p2.circuit]
			currentCircuit.points = append(currentCircuit.points, e.p1)

			p.circuits[e.p2.circuit] = currentCircuit
		} else {
			e.p2.circuit = e.p1.circuit

			currentCircuit := p.circuits[e.p1.circuit]
			currentCircuit.points = append(currentCircuit.points, e.p2)

			p.circuits[e.p1.circuit] = currentCircuit
		}
		return
	}

	// Both points have a circuit (and it's not the same one) - merge them into a new one
	if e.p1.circuit != 0 && e.p2.circuit != 0 && e.p1.circuit != e.p2.circuit {
		c1 := e.p1.circuit
		c2 := e.p2.circuit

		// Make the circuit
		newCircuit := circuit{
			id:     p.nextAvailable,
			points: make([]*point, 0),
		}

		// Set the circuit on all the points of the previous two circuits and add them to this one
		for _, pt := range p.circuits[c1].points {
			pt.circuit = p.nextAvailable
			newCircuit.points = append(newCircuit.points, pt)
		}

		for _, pt := range p.circuits[c2].points {
			pt.circuit = p.nextAvailable
			newCircuit.points = append(newCircuit.points, pt)
		}

		// Delete the old circuits
		delete(p.circuits, c1)
		delete(p.circuits, c2)

		p.circuits[p.nextAvailable] = newCircuit
		p.nextAvailable++
	}
}

func threeLargest(circuits map[int]circuit) int {
	out := 1
	sizes := make([]int, 0, len(circuits))

	for _, circ := range circuits {
		sizes = append(sizes, len(circ.points))
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	for i := 0; i < 3; i++ {
		out *= sizes[i]
	}

	return out
}

func (p playground) allConnected(points []*point) bool {
	if len(p.circuits) != 1 {
		return false
	}

	for _, v := range p.circuits { // should only be one
		if len(v.points) == len(points) {
			return true
		}
	}

	return false
}
