# arangodb-test

## cpu
AMD Ryzen 3 3200U with Radeon Vega Mobile Gfx  

## benchmarks
### binary tree
100k users traversal

getting only vertices
```
BenchmarkTreversalFrom16To16Lvl-4              2         852780368 ns/op // 0,85278s
BenchmarkTreversalFrom13To16Lvl-4              1        2280379442 ns/op // 2,28038s
BenchmarkTreversalFrom10To13Lvl-4              3         381193014 ns/op // 0,38119s
BenchmarkTreversalFrom7To10Lvl-4              27          57173000 ns/op // 0,05717s
BenchmarkTreversalFrom1To16Lvl-4               1        2391857193 ns/op // 2,39186s
```
getting only edges
```
BenchmarkTreversalFrom16To16Lvl-4              2         972496528 ns/op // 0,97250s
BenchmarkTreversalFrom13To16Lvl-4              1        2270353583 ns/op // 2,27035s
BenchmarkTreversalFrom10To13Lvl-4              4         344779165 ns/op // 0,34478s
BenchmarkTreversalFrom7To10Lvl-4              26          53424872 ns/op // 0,05342s
BenchmarkTreversalFrom1To16Lvl-4               1        2409634405 ns/op // 2,40963s
```

getting vertices with edges
```
BenchmarkTreversalFrom16To16Lvl-4              1        1250541776 ns/op // 1,25054s
BenchmarkTreversalFrom13To16Lvl-4              1        3715773536 ns/op // 3,71577s
BenchmarkTreversalFrom10To13Lvl-4              2         607330364 ns/op // 0,60733s
BenchmarkTreversalFrom7To10Lvl-4              20          72164278 ns/op // 0,07216s
BenchmarkTreversalFrom1To16Lvl-4               1        4024310004 ns/op // 4,02431s
```

getting the path to each vertex
```
BenchmarkTreversalFrom16To16Lvl-4              1        12915378680 ns/op // 12,91538s
BenchmarkTreversalFrom13To16Lvl-4              1        33754724155 ns/op // 33,75472s
BenchmarkTreversalFrom10To13Lvl-4              1         4873615153 ns/op // 4,873615s
BenchmarkTreversalFrom7To10Lvl-4               3          489871877 ns/op // 0,489872s
BenchmarkTreversalFrom1To16Lvl-4               1        34967829269 ns/op // 34,96783s
```