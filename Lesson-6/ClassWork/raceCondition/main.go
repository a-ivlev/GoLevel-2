// Программа будет корректно работать без WaitGroup,
// если запустить её на 1 процессоре, следующей командой:
// GOMAXPROCS=1 go run main.go
// Трассировку работы планировщика можно посмотреть следующей командой:
// GODEBUG=schedtrace=100 GOMAXPROCS=2 go run ClassWork/raceCondition/main.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаём WaitGroup
	wg := &sync.WaitGroup{}

	// Инкрементируем счётчик WaitGroup на 2,
	// т.к. будем запускать две горутины.
	wg.Add(2)
	go func() {
		// После завершения работы данной горутины,
		// декрементируем счётчик WaitGroup.
		defer wg.Done()
		for c := 'a'; c <= 'z'; c += 1 {
			fmt.Printf("%c", c)
		}
	}()

	go func() {
		// После завершения работы данной горутины,
		// декрементируем счётчик WaitGroup.
		defer wg.Done()
		for i := 0; i < 10; i += 1 {
			fmt.Printf("%d", i)
		}
	}()

	// Ждём завершения всех горутин.
	wg.Wait()
}