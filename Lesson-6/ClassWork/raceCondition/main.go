// Программа будет корректно работать без WaitGroup,
// если запустить её на 1 процессоре, следующей командой:
// GOMAXPROCS=1 go run main.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаём WaitGroup
	wg := *&sync.WaitGroup{}

	// Инкрементируем значение счётчика WaitGroup на 1,
	// т.к. запускается 1 горутина.
	wg.Add(1)
	go func() {
		defer wg.Done()
		for c := 'a'; c <= 'z'; c += 1 {
			fmt.Printf("%c\n", c)
		}

	}()

	// Ещё раз инкриментируем значение счётчика WaitGroup на 1,
	// т.к. запускается ещё одна горутина.
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i += 1 {
			fmt.Printf("%d\n", i)
		}
	}()

	// Ждём завершения всех горутин.
	wg.Wait()
}