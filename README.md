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
Benchmark_duckmap_parallel_write-4   	 3000000	       446 ns/op	     116 B/op	       2 allocs/op
Benchmark_duckmap_parallel_read-4    	10000000	       155 ns/op	       8 B/op	       0 allocs/op
Benchmark_map_parallel_write-4       	 2000000	       803 ns/op	     175 B/op	       2 allocs/op
Benchmark_map_parallel_read-4        	10000000	       145 ns/op	       0 B/op	       0 allocs/op
```
