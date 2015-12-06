package skylines

import "testing"

// Provide a single building and show
// that the correct critical points are returned.
func testSimple(solver_func Solver, t *testing.T) {

	height := 5
	width := 4
	buildings := []Building{{0, width, height}}

	points := solver_func(buildings)

	if len(points) != 2 {
		t.Errorf("Expected 2 Critical Points. Found: %s", points)
	}

	p1 := CriticalPoint{0, height}
	if !p1.Equals(points[0]) {
		t.Errorf("Error with 1st Point: %s. expected: %s.", points[0], p1)
	}

	p2 := CriticalPoint{width, 0}
	if !p2.Equals(points[1]) {
		t.Errorf("Error with 2nd Point: %s. expected: %s.", points[1], p2)
	}
}

// Provide a set of buildings in an unsorted order
// and show that the correct critical points are found
// and returned in sorted order.
func testHard(solver_func Solver, t *testing.T) {

	expected := []CriticalPoint{
		{2, 10},
		{3, 15},
		{7, 12},
		{12, 0},
		{15, 10},
		{20, 8},
		{24, 0}}
	buildings := []Building{
		{5, 12, 12},
		{2, 9, 10},
		{3, 7, 15},
		{19, 24, 8},
		{15, 20, 10}}

	points := solver_func(buildings)

	if len(points) != len(expected) {
		t.Errorf("Expected %d Critical Points. Found: %d",
			len(expected), len(points))
	}

	for i, p1 := range expected {
		p2 := points[i]
		if !p1.Equals(p2) {
			t.Errorf("Error at index $d expected: %s. found: %s.", p1, p2)
		}
	}
}

func TestNaiveSimple(t *testing.T) {

	testSimple(SolveNaive, t)

}

func TestNaiveHard(t *testing.T) {

	testHard(SolveNaive, t)

}
