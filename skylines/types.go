package skylines

import "fmt"

// A Building is a rectangle on the skyline grounded at 0 elevation.
type Building struct {
	Left   int
	Right  int
	Height int
}

// type used to enable sorting a set of buildings
type Buildings []Building

func (b Buildings) Len() int {
	return len(b)
}

// Return true if the left edge of building at index i
// comes before the left edge of building at index j
func (b Buildings) Less(i, j int) bool {
	return b[i].Left < b[j].Left
}
func (b Buildings) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// A CriticalPoint is a region on the skyline where the height changes.
type CriticalPoint struct {
	X int
	Y int
}

type CriticalPoints []*CriticalPoint

func (c CriticalPoints) Len() int {
	return len(c)
}
func (c CriticalPoints) Less(i, j int) bool {
	// use greater to form a priority queue
	return c[i].Y > c[j].Y
}
func (c CriticalPoints) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c *CriticalPoints) Push(x interface{}) {
	item := x.(*CriticalPoint)
	*c = append(*c, item)
}

func (c *CriticalPoints) Pop() interface{} {
	old := *c
	n := len(old)
	item := old[n-1]
	*c = old[0 : n-1]
	return item
}

func (c *CriticalPoints) Peak() interface{} {
	old := *c
	return old[0]
}

func (c *CriticalPoints) Max() int {
	if len(*c) == 0 {
		return 0
	}
	return c.Peak().(*CriticalPoint).Y
}

// Function Signature for a function that can solve the skylines problem
type Solver func(Buildings) []CriticalPoint

func (p CriticalPoint) String() string {
	return fmt.Sprintf("CriticalPoint(%d,%d)", p.X, p.Y)
}

func (p1 CriticalPoint) Equals(p2 CriticalPoint) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}
