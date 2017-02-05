package duckmap

import (
	"math/rand"
	"sync"
	"time"
	_ "unsafe"
)

const (
	concurrency          = 1024
	duckilingDefaultSize = 256
)

type duckling map[interface{}]interface{}

// Map is awesome lockfree hashtable
type Map struct {
	ducklings []duckling
	mutices   []sync.RWMutex
	seed      uintptr
}

// NewMap is a constructor for the Map
func NewMap() *Map {
	rand.Seed(time.Now().UTC().UnixNano())
	m := &Map{
		seed:      uintptr(rand.Int63()),
		mutices:   make([]sync.RWMutex, concurrency),
		ducklings: make([]duckling, concurrency),
	}
	for b := 0; b < concurrency; b++ {
		m.ducklings[b] = make(duckling, duckilingDefaultSize)
	}
	return m
}

//go:linkname efaceHash runtime.efaceHash
func efaceHash(i interface{}, seed uintptr) uintptr

func (m *Map) bucket(value interface{}) uint64 {
	hash := uint64(efaceHash(value, m.seed))
	return hash % uint64(concurrency)
}
