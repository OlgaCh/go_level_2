package main

import (
	"math/rand"
	"sync"
	"testing"
)

type Set struct {
	sync.Mutex
	mm map[int]struct{}
}

func NewSet() *Set {
	return &Set{
		mm: map[int]struct{}{},
	}
}

func (s *Set) Add(i int) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}

func (s *Set) Has(i int) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.mm[i]
	return ok
}

type RWSet struct {
	sync.RWMutex
	mm map[int]struct{}
}

func NewRWSet() *RWSet {
	return &RWSet{
		mm: map[int]struct{}{},
	}
}

func (s *RWSet) Add(i int) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}

func (s *RWSet) Has(i int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.mm[i]
	return ok
}

func BenchmarkMutex_90_10(b *testing.B) {
	var set = NewSet()

	b.Run("Mutex 90% write 10% read", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Float32() <= 0.1 {
					set.Has(1)
				} else {
					set.Add(1)
				}
			}
		})
	})
}

func BenchmarkRWMutex_90_10(b *testing.B) {
	var set = NewRWSet()

	b.Run("RWMutex 90% write 10% read", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Float32() <= 0.1 {
					set.Has(1)
				} else {
					set.Add(1)
				}
			}
		})
	})
}

func BenchmarkMutex_50_50(b *testing.B) {
	var set = NewSet()

	b.Run("Mutex 50% write 50% read", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Float32() <= 0.5 {
					set.Add(1)
				} else {
					set.Has(1)
				}
			}
		})
	})
}

func BenchmarkRWMutex_50_50(b *testing.B) {
	var set = NewRWSet()

	b.Run("RWMutex 50% write 50% read", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Float32() <= 0.5 {
					set.Add(1)
				} else {
					set.Has(1)
				}
			}
		})
	})
}

func BenchmarkMutex_10_90(b *testing.B) {
	var set = NewSet()

	b.Run("Mutex 10% write 90% read", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Float32() <= 0.1 {
					set.Add(1)
				} else {
					set.Has(1)
				}
			}
		})
	})
}

func BenchmarkRWMutex_10_90(b *testing.B) {
	var set = NewRWSet()

	b.Run("RWMutex 10% write 90% read", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Float32() <= 0.1 {
					set.Add(1)
				} else {
					set.Has(1)
				}
			}
		})
	})
}
