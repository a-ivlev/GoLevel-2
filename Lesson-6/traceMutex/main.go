// Написать программу, которая использует мьютекс для безопасного доступа к данным
// из нескольких потоков. Выполнить трассировку программы.
// A app that uses a mutex to securely access data.

// Трассировку работы программы можно посмотреть следующей командой:
// GODEBUG=schedtrace=100 go run traceMutex/main.go

// Или включить в код приложения и запустить на исполнение следующей командой:
// GOMAXPROCS=1 go run traceMutex/main.go 2>traceMutex/trace.out
// Посмотреть результат работы можно набрав в консоле:
// go tool trace traceMutex/trace.out

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
	trace.Start(os.Stderr)
	defer trace.Stop()

	count := 0
	num := 10

	// В данном случае создаём переменные mut и wgне по указателю,
	// т.к. в фукции замыкания захват всегда происходит по указателю.
	mut := sync.Mutex{}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < num; i++ {
			mut.Lock()
			count += i
			mut.Unlock()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < num; i++ {
			mut.Lock()
			count += i
			mut.Unlock()
		}
	}()

	wg.Wait()

	fmt.Println(count)

}
