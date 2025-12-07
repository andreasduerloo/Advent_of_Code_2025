package day07

type set map[int]struct{}

func countSplits(diagram []string) int {
	var out int
	beams := make(map[int]struct{})

	for r, row := range diagram {
		if r == 0 { // find out where the beam starts
			for c, char := range row {
				if char == 'S' {
					beams[c] = struct{}{}
					break
				}
			}
		} else {
			// scan through the rows, if there is a '^' in a beam column, split the beam and remove the old one
			for c, char := range row {
				if char == '^' {
					if _, present := beams[c]; present {
						delete(beams, c)
						out++
						beams[c-1] = struct{}{}
						beams[c+1] = struct{}{}
					}
				}
			}
		}

	}

	return out
}

func totalPaths(diagram []string) int {
	paths := make(map[int]int)

	for r, row := range diagram {
		if r == 0 { // find out where the beam starts
			for c, char := range row {
				if char == 'S' {
					paths[c] = 1
					break
				}
			}
		} else {
			// scan through the rows, if there is a '^' in a beam column, split the beam and remove the old one
			for c, char := range row {
				if char == '^' {
					if count, present := paths[c]; present {
						delete(paths, c)
						paths[c-1] += count
						paths[c+1] += count
					}
				}
			}
		}
	}

	var out int
	for _, v := range paths {
		out += v
	}

	return out
}
