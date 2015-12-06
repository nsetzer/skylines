package skylines

import (
	"container/heap"
	"sort"
	"testing"
)

// Show that buildings can be sorted by left edge
func TestSortBuildings(t *testing.T) {
	buildings := Buildings{
		{5, 12, 12},
		{2, 9, 10},
		{3, 7, 15}}

	sort.Sort(buildings)

	expected_order := []int{2, 3, 5}

	for i, v := range expected_order {
		if buildings[i].Left != v {
			t.Errorf("call to sort failed at index %d.", i)
		}
	}
}

func testHeapHeight(t *testing.T, hq *CriticalPoints, height int) {
	h := hq.Max()
	if h != height {
		t.Fatalf("heap height error. Expected: %d. Found: %d.", height, h)
	}
}

func TestCriticalPointMaxHeap(t *testing.T) {

	// contruct a new heap
	hq := make(CriticalPoints, 0)

	// an empty heap has height 0
	testHeapHeight(t, &hq, 0)

	heap.Push(&hq, &CriticalPoint{X: 7, Y: 5})
	testHeapHeight(t, &hq, 5)

	// add an item smaller than current maximum
	heap.Push(&hq, &CriticalPoint{X: 5, Y: 3})
	testHeapHeight(t, &hq, 5)

	// add an item larger than current maximum
	heap.Push(&hq, &CriticalPoint{X: 3, Y: 7})
	testHeapHeight(t, &hq, 7)

	heap.Pop(&hq)
	testHeapHeight(t, &hq, 5)

	heap.Pop(&hq)
	testHeapHeight(t, &hq, 3)

	heap.Pop(&hq)
	testHeapHeight(t, &hq, 0)

}

// Provide a single building and show
// that the correct critical points are returned.
func testSimple(solver_func Solver, t *testing.T) {

	height := 5
	width := 4
	buildings := Buildings{{0, width, height}}

	points := solver_func(buildings)

	if len(points) != 2 {
		t.Fatalf("Expected 2 Critical Points. Found: %s", points)
	}

	p1 := CriticalPoint{X: 0, Y: height}
	if !p1.Equals(points[0]) {
		t.Fatalf("Error with 1st Point: %s. expected: %s.", points[0], p1)
	}

	p2 := CriticalPoint{X: width, Y: 0}
	if !p2.Equals(points[1]) {
		t.Fatalf("Error with 2nd Point: %s. expected: %s.", points[1], p2)
	}
}

// Provide a set of buildings in an unsorted order
// and show that the correct critical points are found
// and returned in sorted order.
func testHard(solver_func Solver, t *testing.T) {

	expected := []CriticalPoint{
		{X: 2, Y: 10},
		{X: 3, Y: 15},
		{X: 7, Y: 12},
		{X: 12, Y: 0},
		{X: 15, Y: 10},
		{X: 20, Y: 8},
		{X: 24, Y: 0}}
	buildings := Buildings{
		{5, 12, 12},
		{2, 9, 10},
		{3, 7, 15},
		{19, 24, 8},
		{15, 20, 10}}

	points := solver_func(buildings)

	if len(points) != len(expected) {
		t.Fatalf("Expected %d Critical Points. Found: %d",
			len(expected), len(points))
	}

	for i, p1 := range expected {
		p2 := points[i]
		if !p1.Equals(p2) {
			t.Fatalf("Error at index %d expected: %s. found: %s.", i, p1, p2)
		}
	}
}

func TestNaiveSimple(t *testing.T) {
	testSimple(SolveNaive, t)
}

func TestNaiveHard(t *testing.T) {
	testHard(SolveNaive, t)
}

func TestFastSimple(t *testing.T) {
	testSimple(SolveFast, t)
}

func TestFastHard(t *testing.T) {
	testHard(SolveFast, t)
}
