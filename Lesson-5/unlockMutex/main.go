package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var number int64
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			// Критическая секция
			number++
			// Выводя на печать, можно увидеть разницу работу программы с mutex и без mutex.
			fmt.Printf("Число number = %d изменено i = %d горутиной.\n", number, i)
			// Sleep имитирует задержку выполнения программы.
			time.Sleep(time.Millisecond * 100)
			// После задержки выводим на печать текущее значение number и номер горутины.
			fmt.Printf("Число number = %d изменено i = %d горутиной.\n", number, i)
		}(i)
	}
	wg.Wait()
	fmt.Println("number = ", number)
}
