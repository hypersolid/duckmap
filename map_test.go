package duckmap

import (
	"math/rand"
	"regexp"
	"runtime"
	"sync"
	"testing"
)

const (
	testRange = 1000
	testProcs = 3
)

func Test_Set_works(t *testing.T) {
	m := NewMap()
	k := rand.Int()
	v := rand.Int()
	m.Set(k, v)
	if m.Get(k) != v {
		t.Error("not set", k, v, m.Get(k))
	}
}

func Test_Get_works(t *testing.T) {
	m := NewMap()
	k := rand.Int()
	v := rand.Int()
	m.Set(k, v)
	if m.Get(k) != v {
		t.Error("incorrect get", k, v, m.Get(k))
	}
}

func Test_Delete_works(t *testing.T) {
	m := NewMap()
	k := rand.Int()
	v := rand.Int()
	m.Set(k, v)
	m.Delete(k)
	if m.Get(k) != nil {
		t.Error("not deleted", k, v, m.Get(k))
	}
}

func Test_Keys_works(t *testing.T) {
	m := NewMap()
	checksum1 := 0
	for i := 0; i < testRange; i++ {
		m.Set(i, i^2)
		checksum1 += i
	}

	keys := m.Keys()
	checksum2 := 0
	for i := 0; i < testRange; i++ {
		checksum2 += keys[i].(int)
	}

	if checksum1 != checksum2 {
		t.Error("key checksum error", checksum1, checksum2)
	}
}

func Test_Values_works(t *testing.T) {
	m := NewMap()
	checksum1 := 0
	for i := 0; i < testRange; i++ {
		m.Set(i^2, i)
		checksum1 += i
	}

	values := m.Values()
	checksum2 := 0
	for i := 0; i < testRange; i++ {
		checksum2 += values[i].(int)
	}

	if checksum1 != checksum2 {
		t.Error("key checksum error", checksum1, checksum2)
	}
}

func Test_String_works(t *testing.T) {
	m := NewMap()
	m.Set("test", true)
	match, _ := regexp.MatchString(`DuckMap<\d+>`, m.String())
	if !match {
		t.Error("to string conversion error", m.String())
	}
}

func Test_parallel_operation(t *testing.T) {
	procs := runtime.GOMAXPROCS(testProcs)
	m := NewMap()
	var wg sync.WaitGroup
	for i := 0; i < testProcs; i++ {
		wg.Add(1)
		go func() {
			for k := 0; k < testRange*testRange; k++ {
				v := rand.Intn(testRange)
				m.Set(v, v)
				m.Get(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	runtime.GOMAXPROCS(procs)
}

// benchmarks
func Benchmark_duckmap_parallel_write(b *testing.B) {
	b.StopTimer()
	procs := runtime.GOMAXPROCS(testProcs)
	start := make(chan struct{}, testProcs)
	done := make(chan struct{}, testProcs)
	m := NewMap()
	for k := 0; k < testProcs; k++ {
		go func() {
			<-start
			for i := 0; i < b.N/testProcs; i++ {
				m.Set(rand.Int(), rand.Int())
			}
			done <- struct{}{}
		}()
	}
	b.StartTimer()

	for k := 0; k < testProcs; k++ {
		start <- struct{}{}
	}
	for k := 0; k < testProcs; k++ {
		<-done
	}
	runtime.GOMAXPROCS(procs)
}

func Benchmark_duckmap_parallel_read(b *testing.B) {
	b.StopTimer()
	procs := runtime.GOMAXPROCS(testProcs)
	start := make(chan struct{}, testProcs)
	done := make(chan struct{}, testProcs)
	m := NewMap()
	for k := 0; k < testProcs; k++ {
		go func() {
			<-start
			for i := 0; i < b.N/testProcs; i++ {
				m.Get(rand.Intn(testRange))
			}
			done <- struct{}{}
		}()
	}
	for k := 0; k < testRange; k++ {
		m.Set(k, k)
	}
	b.StartTimer()

	for k := 0; k < testProcs; k++ {
		start <- struct{}{}
	}
	for k := 0; k < testProcs; k++ {
		<-done
	}
	runtime.GOMAXPROCS(procs)
}

func Benchmark_map_parallel_write(b *testing.B) {
	b.StopTimer()
	procs := runtime.GOMAXPROCS(testProcs)
	start := make(chan struct{}, testProcs)
	done := make(chan struct{}, testProcs)
	m := make(map[interface{}]interface{})
	mtx := new(sync.RWMutex)
	for k := 0; k < testProcs; k++ {
		go func() {
			<-start
			for i := 0; i < b.N/testProcs; i++ {
				mtx.Lock()
				m[rand.Int()] = rand.Int()
				mtx.Unlock()
			}
			done <- struct{}{}
		}()
	}
	b.StartTimer()

	for k := 0; k < testProcs; k++ {
		start <- struct{}{}
	}
	for k := 0; k < testProcs; k++ {
		<-done
	}
	runtime.GOMAXPROCS(procs)
}

func Benchmark_map_parallel_read(b *testing.B) {
	b.StopTimer()
	procs := runtime.GOMAXPROCS(testProcs)
	start := make(chan struct{}, testProcs)
	done := make(chan struct{}, testProcs)
	m := make(map[interface{}]interface{})
	mtx := new(sync.RWMutex)
	var wasted int
	for k := 0; k < testProcs; k++ {
		go func() {
			<-start
			for i := 0; i < b.N/testProcs; i++ {
				mtx.RLock()
				wasted = m[rand.Intn(testRange)].(int)
				mtx.RUnlock()
			}
			done <- struct{}{}
		}()
	}
	for k := 0; k < testRange; k++ {
		m[k] = k
	}
	b.StartTimer()

	for k := 0; k < testProcs; k++ {
		start <- struct{}{}
	}
	for k := 0; k < testProcs; k++ {
		<-done
	}
	runtime.GOMAXPROCS(procs)
}
