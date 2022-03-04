package main

import (
	"fmt"
	"sync"
)

const count = 1000

func main() {
	var (
		counter int
		// Создаем экземпляр
		wg = sync.WaitGroup{}
	)

	// Инициализируем семафор исходным состоянием
	wg.Add(count)
	for i := 0; i < count; i += 1 {
		go func() {
			counter += 1
			// Выполняем декремент семафора
			wg.Done()
		}()
	}
	// Ждем обнуления семафора
	wg.Wait()
	// Выводим показание общего счётчика
	fmt.Println(counter)
}

// package main

// import (
// 	"sync"
// 	"testing"
// )

// type Set struct {
// 	sync.Mutex
// 	mm map[int]struct{}
// }

// func NewSet() *Set {
// 	return &Set{
// 		mm: map[int]struct{}{},
// 	}
// }
// func (s *Set) Add(i int) {
// 	s.Lock()
// 	s.mm[i] = struct{}{}
// 	s.Unlock()
// }
// func (s *Set) Has(i int) bool {
// 	s.Lock()
// 	defer s.Unlock()
// 	_, ok := s.mm[i]
// 	return ok
// }

// func BenchmarkSetAdd(b *testing.B) {
// 	var set = NewSet()
// 	b.Run("", func(b *testing.B) {
// 		b.SetParallelism(1000)
// 		b.RunParallel(func(pb *testing.PB) {
// 			for pb.Next() {
// 				set.Add(1)
// 			}
// 		})
// 	})
// }

// func BenchmarkSetHas(b *testing.B) {
// 	var set = NewSet()
// 	b.Run("", func(b *testing.B) {
// 		b.SetParallelism(1000)
// 		b.RunParallel(func(pb *testing.PB) {
// 			for pb.Next() {
// 				set.Has(1)
// 			}
// 		})
// 	})
// }
