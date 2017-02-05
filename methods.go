package duckmap

import (
	"fmt"
	"unsafe"
)

// Get reads entry
func (m *Map) Get(key interface{}) (result interface{}) {
	b := m.bucket(key)
	m.mutices[b].RLock()
	result = m.ducklings[b][key]
	m.mutices[b].RUnlock()
	return
}

// Set adds or replaces entry
func (m *Map) Set(key, value interface{}) {
	b := m.bucket(key)
	m.mutices[b].Lock()
	m.ducklings[b][key] = value
	m.mutices[b].Unlock()
}

// Delete removes entry
func (m *Map) Delete(key interface{}) {
	b := m.bucket(key)
	m.mutices[b].Lock()
	delete(m.ducklings[b], key)
	m.mutices[b].Unlock()
}

// Keys returns map keys
func (m *Map) Keys() []interface{} {
	result := make([]interface{}, 0)
	for b := 0; b < m.concurrency; b++ {
		m.mutices[b].RLock()
		internalResult := make([]interface{}, len(m.ducklings[b]))
		i := 0
		for k := range m.ducklings[b] {
			internalResult[i] = k
			i++
		}
		m.mutices[b].RUnlock()
		result = append(result, internalResult...)
	}
	return result
}

// Values returns map values
func (m *Map) Values() []interface{} {
	result := make([]interface{}, 0)
	for b := 0; b < m.concurrency; b++ {
		m.mutices[b].RLock()
		internalResult := make([]interface{}, len(m.ducklings[b]))
		i := 0
		for _, v := range m.ducklings[b] {
			internalResult[i] = v
			i++
		}
		m.mutices[b].RUnlock()
		result = append(result, internalResult...)
	}
	return result
}

func (m *Map) String() string {
	return fmt.Sprintf(
		"DuckMap<%d>",
		unsafe.Pointer(m),
	)
}
