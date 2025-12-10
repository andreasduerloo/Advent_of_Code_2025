package day09

import "math"

type point struct {
	x int
	y int
}

func toPoint(coords []int) point {
	return point{
		x: coords[0],
		y: coords[1],
	}
}

func convert(coords [][]int) []point {
	out := make([]point, 0, len(coords))

	for _, pair := range coords {
		out = append(out, toPoint(pair))
	}

	return out
}

func area(p1, p2 point) int {
	return int(math.Abs(float64(p1.x-p2.x)+1) * math.Abs(float64(p1.y-p2.y)+1))
}

func largestArea(points []point) int {
	var largest int

	for i := range len(points) - 2 {
		for j := i + 1; j < len(points); j++ {
			if area(points[i], points[j]) > largest {
				largest = area(points[i], points[j])
			}
		}
	}

	return largest
}
