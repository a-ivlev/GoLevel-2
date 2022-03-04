/*
Задание 2
Написать многопоточную программу, в которой будет использоваться явный вызов планировщика.
Выполните трассировку программы
*/

/*
Трассировку работы программы можно посмотреть следующей командой:
GODEBUG=schedtrace=100 go run main.go

Или включить в код приложения и запустить на исполнение следующей командой:
go run ./main.go 2> ./trace.out

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

	count := 10
	chIn := make(chan int)
	resault := make(chan int)
	wg := sync.WaitGroup{}

	// В горутине генерируем данные.
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(chIn)
		for i := 1; i <= count; i++ {
			chIn <- i * 2
		}
	}()

	// В другой горутине производим обработку данных.
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(resault)
		for elem := range chIn {
			resault <- elem * elem
		}
	}()

	// В главной горутине выводим результат на печать.
	for elem := range resault {
		fmt.Println(elem)
	}

	wg.Wait()
}
