# skylines
implementation of skylines problem in golang

https://leetcode.com/problems/the-skyline-problem/

#### Benchmark Results:
benchmark was run using:
```
    go test -bench=".*" github.com/nsetzer/skylines/skylines
```

go benchmark runs a test with increasing N until a useful timing measurement
is made. For a given N, total runtime is N*(ns/op). The number of buildings
is given below to show what input size was used to produce the measurement.

| Solver Function | Number of Buildings | Time to Complete |
| --------------- | ------------------- | ---------------- |
| SolveNaive      |              100000 |     147261 ns/op |
| SolveFast       |             1000000 |       1085 ns/op |


