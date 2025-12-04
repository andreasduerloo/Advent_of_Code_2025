package day04

func scanForAccess(grid []string) (int, int, []string) {
	var accessible int
	var removed int
	newGrid := []string{}

	maxRow := len(grid) - 1
	maxCol := len(grid[0]) - 1

	for row, r := range grid {
		newLine := ""
		for col, c := range r {
			if rune(c) == '@' {
				var localCount int
				// Up
				if row > 0 {
					// Up + left
					if col > 0 {
						if rune(grid[row-1][col-1]) == '@' {
							localCount++
						}
					}
					// Up
					if rune(grid[row-1][col]) == '@' {
						localCount++
					}
					// Up + right
					if col < maxCol {
						if rune(grid[row-1][col+1]) == '@' {
							localCount++
						}
					}
				}

				// Down
				if row < maxRow {
					// Down + left
					if col > 0 {
						if rune(grid[row+1][col-1]) == '@' {
							localCount++
						}
					}
					// Down
					if rune(grid[row+1][col]) == '@' {
						localCount++
					}
					// Down + right
					if col < maxCol {
						if rune(grid[row+1][col+1]) == '@' {
							localCount++
						}
					}
				}

				// Left and right
				if col > 0 {
					if rune(grid[row][col-1]) == '@' {
						localCount++
					}
				}
				if col < maxCol {
					if rune(grid[row][col+1]) == '@' {
						localCount++
					}
				}

				if localCount < 4 {
					accessible++
					newLine += "."
					removed++
				} else {
					newLine += "@"
				}
			} else {
				newLine += "."
			}
		}
		newGrid = append(newGrid, newLine)
	}

	return accessible, removed, newGrid
}

func reduce(grid []string) int {
	_, removed, newGrid := scanForAccess(grid)

	_, afterPass, newGrid := scanForAccess(newGrid)

	for afterPass != 0 {
		removed += afterPass
		_, afterPass, newGrid = scanForAccess(newGrid)
	}

	return removed
}
