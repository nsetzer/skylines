package skylines

// Solve the skyline problem the naive way.
// O(n^2) in the worst case if every building
// is the same width. It also strongly depends on the
// input. Wider skylines will take longer to process.
func SolveNaive(buildings Buildings) (points []CriticalPoint) {

	//first determine the width of the skyline
	width := 0
	for _, b := range buildings {
		if b.Right > width {
			width = b.Right
		}
	}

	// determine the maximum height at each point
	skyline := make([]int, width+1)
	for _, b := range buildings {
		for i := b.Left; i <= b.Right; i++ {
			if skyline[i] < b.Height {
				skyline[i] = b.Height
			}
		}
	}

	// read off the critical points from the skyline
	last_height := 0
	for i, h := range skyline {
		if h < last_height {
			points = append(points, CriticalPoint{X:i - 1, Y:h})
		} else if h > last_height {
			points = append(points, CriticalPoint{X:i, Y:h})
		}
		last_height = h
	}
	points = append(points, CriticalPoint{X:width, Y:0})
	return
}
