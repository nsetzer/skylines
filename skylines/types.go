package skylines

import "fmt"

// A Building is a rectangle on the skyline grounded at 0 elevation.
type Building struct {
	Left   int
	Right  int
	Height int
}

// A CriticalPoint is a region on the skyline where the height changes.
type CriticalPoint struct {
	X int
	Y int
}

// Function Signature for a funciton that can solve the skylines problem
type Solver func([]Building) []CriticalPoint

func (p CriticalPoint) String() string {
	return fmt.Sprintf("CriticalPoint(%d,%d)", p.X, p.Y)
}

func (p1 CriticalPoint) Equals(p2 CriticalPoint) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}
