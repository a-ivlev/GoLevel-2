// Бенчмарк запускается следующей командой:
// go test -bench=. mainCh_test.go

// Запуск бенчмарка в режиме параллельного выполнения:
// go test -parallel=8 -bench=. mainCh_test.go
// разобраться почему в параллельном режиме не работает!

package main

import (
	// "sync"
	"testing"
)

type Set struct {
	//sync.Mutex
	chMutex chan struct{}
	mm map[int]struct{}
}

func NewSet() *Set {
	return &Set{
		chMutex: make(chan struct{}, 1),
		mm: map[int]struct{}{},
	}
}
func (s *Set) Add(i int) {
	//s.Lock()
	s.chMutex <- struct{}{}
	s.mm[i] = struct{}{}
	//s.Unlock()
	<- s.chMutex
}
func (s *Set) Has(i int) bool {
	// s.Lock()
	// defer s.Unlock()
	s.chMutex <- struct{}{}
	_, ok := s.mm[i]
	<- s.chMutex
	return ok
}

func (s *Set) Close() {
	close(s.chMutex) 
}

func BenchmarkSetAdd(b *testing.B) {
	var set = NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1)
			}
		})
	})
	set.Close()
}

func BenchmarkSetHas(b *testing.B) {
	var set = NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
	set.Close()
}
