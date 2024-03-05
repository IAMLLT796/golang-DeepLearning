package main

import "time"

type MyConcurrentMap struct {
}

func (m *MyConcurrentMap) Put(k, v int) {

}

func (m *MyConcurrentMap) Get(k int, maxWaitingDuration time.Duration) (int, error) {

}
