# skylines
implementation of skylines problem in golang

#### Benchmark Results:
benchmark was run using
```
    go test -bench=.
```
| Solver Function | Number of Buildings | Time to Complete |
| --------------- | ------------------- | ---------------- |
| SolveNaive      |              100000 |     147883 ns/op |
| SolveFast       |             1000000 |       1101 ns/op |
