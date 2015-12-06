package skylines

import (
	"testing"
	"math/rand"
)

// Randomly generate a set of buildings
func init_benchmark(N int) Buildings {
	buildings := make(Buildings,N)

	rand.Seed(4)

	building_min_height := 1
	building_max_height := 100
	skyline_width := (1+N)*10
	building_max_width := skyline_width/2

	for i:=0; i<N; i++ {
		// determine the width of the building, then
		// place it on somewhere on the skyline where it will fit
		w := rand.Intn( building_max_width )
		buildings[i].Left = rand.Intn( skyline_width - w )
		buildings[i].Right = buildings[i].Left + w
		buildings[i].Height = building_min_height + rand.Intn( building_max_height )
	}

	return buildings
}

func BenchmarkSolveNaive(b *testing.B) {

	buildings := init_benchmark( b.N )
	b.ResetTimer()
	SolveNaive(buildings)
}

func BenchmarkSolveFast(b *testing.B) {

	buildings := init_benchmark( b.N )
	b.ResetTimer()
	SolveFast(buildings)
}