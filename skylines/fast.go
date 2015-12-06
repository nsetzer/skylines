package skylines

import (
	"container/heap"
	"sort"
	//"fmt"
)

func SolveFast(buildings Buildings) (points []CriticalPoint) {

	sort.Sort(buildings)
	hq := make(CriticalPoints, 0)

	idx := 0
	for _, b := range buildings {
		for len(hq) > 0 {
			p := hq.Peak().(*CriticalPoint)
			if p.X < b.Left {
				heap.Pop(&hq)
				// this point cannot contribute to a critical point
				if (p.X < idx) {
					continue
				}
				idx = p.X
				for len(hq)>0 && hq.Peak().(*CriticalPoint).X <= idx {
					heap.Pop(&hq)
				}
				points = append(points, CriticalPoint{X: p.X, Y: hq.Max()})
			} else {
				break;
			}
		}

		h := hq.Max()
		heap.Push(&hq, &CriticalPoint{X: b.Right, Y: b.Height})
		idx = b.Left
		if b.Height > h {
			points = append(points, CriticalPoint{X: b.Left, Y: b.Height})
		}
	}

	for len(hq)>0 {
		p := heap.Pop(&hq).(*CriticalPoint)
		for len(hq)>0 && hq.Peak().(*CriticalPoint).X <= idx {
			heap.Pop(&hq)
		}
		points = append(points, CriticalPoint{X: p.X, Y: hq.Max()})
	}

	return
}
