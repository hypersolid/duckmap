# Duckmap
## Concurrent and DRY interface{} to interface{} map

### Quickstart
```
m := duckmap.NewMap()

m.Set(4, "this")
m.Set(5, "that")

m.Delete(5)

fmt.Println(4, m.Get(4).(string))
fmt.Println(5, m.Get(5))
```

### Benchmarks
```
Benchmark_duckmap_parallel_write-4   	20000000	       575 ns/op
Benchmark_duckmap_parallel_read-4    	50000000	       151 ns/op
Benchmark_map_parallel_write-4       	10000000	       956 ns/op
Benchmark_map_parallel_read-4        	50000000	       148 ns/op
```
