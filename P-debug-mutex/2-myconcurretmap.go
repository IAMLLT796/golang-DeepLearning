package main

import "sync"

type MyConcurrentMap struct {
	sync.Mutex
	// sync.RWMutex
	mp map[int]int
}

func NewMyConcurrentMap() *MyConcurrentMap {
	return &MyConcurrentMap{
		mp: make(map[int]int),
	}
}

func (m *MyConcurrentMap) Get(key int) (int, bool) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	val, ok := m.mp[key]

	return val, ok
}

func (m *MyConcurrentMap) Set(key, val int) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	m.mp[key] = val
}

func (m *MyConcurrentMap) Delete(key int) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	delete(m.mp, key)
}
