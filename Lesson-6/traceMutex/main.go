/*
Задание 1
Написать программу, которая использует мьютекс для безопасного доступа к данным из нескольких потоков.
Выполнить трассировку программы.
*/

/*
Трассировку работы программы можно посмотреть следующей командой:
GODEBUG=schedtrace=100 go run main.go

Или включить в код приложения и запустить на исполнение следующей командой:
GOMAXPROCS=8 go run ./main.go 2> ./trace.out

Посмотреть результат работы можно набрав в консоли:
go tool trace ./trace.out
*/

package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

func main() {
	// Трассировку можно включить прямо в коде программы.
	// Функция Start начинает трассировку фрагмента исходного кода до тех пор,
	// пока не будет вызвана функция Stop.
	// Трассировка пишется в экземпляр io.Writer, переданный функции Start параметром.
	err := trace.Start(os.Stderr)
	if err != nil {
		return
	}
	defer trace.Stop()

	var (
		mu = sync.RWMutex{}
		wg = sync.WaitGroup{}
	)

	// set map которую будем заполнять в одной горутине.
	set := map[int]int{}
	// count колличество значений в map.
	count := 10
	// goroutines колличество горутин которые производят чтение из map.
	goroutines := 8

	// Запускаем горутину которая заполняет map значениями.
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < count; i++ {
			mu.Lock()
			set[i] = i
			mu.Unlock()
		}
	}()

	// Производим запуск читающих горутин.
	wg.Add(goroutines)
	for i := 1; i <= goroutines; i++ {
		g := i
		go func(g int) {
			defer wg.Done()
			for i := 0; i < count; i++ {
				// В фукции замыкания захват всегда происходит по указателю.
				mu.RLock()
				fmt.Printf("goroutine %d, map[%d] = %d\n", g, i, set[i])
				mu.RUnlock()
			}
		}(g)
	}

	wg.Wait()
}
