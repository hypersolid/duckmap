# Duckmap
## Concurrent and DRY interface to interface map

### Quickstart
```
m := duckmap.NewMap()

m.Set(4, "this")
m.Set(5, "that")

m.Delete(5)

fmt.Println(m.Keys())
fmt.Println(m.Values())

fmt.Println(4, m.Get(4).(string))
fmt.Println(5, m.Get(5))
```

### Benchmarks
```
Benchmark_duckmap_parallel_write-4   	10000000	       530 ns/op
Benchmark_duckmap_parallel_read-4    	20000000	       156 ns/op
Benchmark_map_parallel_write-4       	 5000000	       762 ns/op
Benchmark_map_parallel_read-4        	20000000	       144 ns/op
```
