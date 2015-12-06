package skylines

import (
	"container/heap"
	"sort"
)

// remove points which can no longer contribute to a
// Critical Point from the Heap
func removeInvalidPoints(hq *CriticalPoints, edge int) {
	for len(*hq) > 0 && hq.Peak().(*CriticalPoint).X <= edge {
		heap.Pop(hq)
	}
}

// Solve the skyline problem using a Max Heap
//
// O(nlogn) to preprocess the input.
// Approximatly O(2n) to process each critical point, 2 per building
func SolveFast(buildings Buildings) (points []CriticalPoint) {

	sort.Sort(buildings)
	hq := make(CriticalPoints, 0)

	for _, b := range buildings {

		// process previous critical points to see if they have terminated
		for len(hq) > 0 && hq.Peak().(*CriticalPoint).X < b.Left {
			p := heap.Pop(&hq).(*CriticalPoint)
			removeInvalidPoints(&hq, p.X)
			points = append(points, CriticalPoint{X: p.X, Y: hq.Max()})
		}

		current_height := hq.Max()
		// add the current building to the heap
		// store the height with the right edge so that we can
		// remove the point from the heap at the correct time.
		heap.Push(&hq, &CriticalPoint{X: b.Right, Y: b.Height})
		removeInvalidPoints(&hq, b.Left) // if any
		if b.Height > current_height {
			points = append(points, CriticalPoint{X: b.Left, Y: b.Height})
		}
	}

	// process any remaining critical points
	for len(hq) > 0 {
		p := heap.Pop(&hq).(*CriticalPoint)
		removeInvalidPoints(&hq, p.X)
		points = append(points, CriticalPoint{X: p.X, Y: hq.Max()})
	}

	return
}
