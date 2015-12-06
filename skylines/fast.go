package skylines

import (
	"container/heap"
	"sort"
)

// remove points which can no longer contribute to a
// Critical Point from the Heap
func removeInvalidPoints(hq *CriticalPoints, edge int) {
	for len(*hq)>0 && hq.Peak().(*CriticalPoint).X <= edge {
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

	idx := 0 // index of the last critical point found
	for _, b := range buildings {
		for len(hq) > 0 {
			p := hq.Peak().(*CriticalPoint)
			if p.X < b.Left {
				heap.Pop(&hq)
				idx = p.X
				removeInvalidPoints(&hq,idx)
				points = append(points, CriticalPoint{X: p.X, Y: hq.Max()})
			} else {
				break;
			}
		}

		current_height := hq.Max()
		idx = b.Left
		heap.Push(&hq, &CriticalPoint{X: b.Right, Y: b.Height})
		removeInvalidPoints(&hq,idx) // if any
		if b.Height > current_height {
			points = append(points, CriticalPoint{X: b.Left, Y: b.Height})
		}
	}

	for len(hq)>0 {
		p := heap.Pop(&hq).(*CriticalPoint)
		removeInvalidPoints(&hq,p.X)
		points = append(points, CriticalPoint{X: p.X, Y: hq.Max()})
	}

	return
}
