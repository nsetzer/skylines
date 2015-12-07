# skylines
implementation of skylines problem in golang

https://leetcode.com/problems/the-skyline-problem/

#### Benchmark Results:
benchmark was run using:
```
    go test -bench=".*" github.com/nsetzer/skylines/skylines
```

| Solver Function | Number of Buildings | Time to Complete |
| --------------- | ------------------- | ---------------- |
| SolveNaive      |              100000 |     147261 ns/op |
| SolveFast       |             1000000 |       1085 ns/op |
