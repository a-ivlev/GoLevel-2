/*
Тестирование производительности операций чтения и записи на множестве действительных чисел,
безопасность которого обеспечивается sync.Mutex и sync.RWMutex для разных вариантов использования:
10% запись, 90% чтение;
50% запись, 50% чтение;
90% запись, 10% чтение.
*/

package main

import (
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"testing"
)

func BenchmarkMutex10Write90Read(b *testing.B) {
	var (
		mu      sync.Mutex
		counter = 1
		set     = make(map[int]int, 1000)
	)
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Float32() < 0.1 {
					mu.Lock()
					counter += 1
					set[counter] = counter
					mu.Unlock()
				}
				mu.Lock()
				_, _ = set[counter]
				mu.Unlock()

			}
		})
	})
}

func BenchmarkMutex50Write50Read(b *testing.B) {
	var (
		mu      sync.Mutex
		counter = 1
		set     = make(map[int]int, 1000)
	)
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Float32() < 0.5 {
					mu.Lock()
					counter += 1
					set[counter] = counter
					mu.Unlock()
				}
				mu.Lock()
				_, _ = set[counter]
				mu.Unlock()
			}
		})
	})
}

func BenchmarkMutex90Write10Read(b *testing.B) {
	var (
		mu      sync.Mutex
		counter = 1
		set     = make(map[int]int, 1000)
	)
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Float32() < 0.9 {
					mu.Lock()
					counter += 1
					set[counter] = counter
					mu.Unlock()
				}
				mu.Lock()
				_, _ = set[counter]
				mu.Unlock()
			}
		})
	})
}

func BenchmarkRWMutex10Write90Read(b *testing.B) {
	var (
		mu      sync.RWMutex
		counter = 1
		set     = make(map[int]int, 1000)
	)
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Float32() < 0.1 {
					mu.Lock()
					counter += 1
					set[counter] = counter
					mu.Unlock()
				}
				mu.RLock()
				_, _ = set[counter]
				mu.RUnlock()
			}
		})
	})
}

func BenchmarkRWMutex50Write50Read(b *testing.B) {
	var (
		mu      sync.RWMutex
		counter = 1
		set     = make(map[int]int, 1000)
	)
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Float32() < 0.5 {
					mu.Lock()
					counter += 1
					set[counter] = counter
					mu.Unlock()
				}
				mu.RLock()
				_, _ = set[counter]
				mu.RUnlock()
			}
		})
	})
}

func BenchmarkRWMutex90Write10Read(b *testing.B) {
	var (
		mu      sync.RWMutex
		counter = 1
		set     = make(map[int]int, 1000)
	)
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Float32() < 0.9 {
					mu.Lock()
					counter += 1
					set[counter] = counter
					mu.Unlock()
				}
				mu.RLock()
				_, _ = set[counter]
				mu.RUnlock()
			}
		})
	})
}
