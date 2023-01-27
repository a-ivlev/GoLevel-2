/*
Выполните сборку ваших предыдущих программ под операционную систему,
отличающуюся от текущей. Проанализируйте вывод команды file для полученного
исполняемого файла. Попробуйте запустить исполняемый файл.
*/

// Программа числовой конвеер.
//
// Использует два канала:
// в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала выводяться в stdout.
package main

import "fmt"

// WriteCh получает на вход массив array и канал chIn. В go-рутине читаем числа из полученного массива array и пишем в канал chIn.
func WriteCh(array []int64, chIn chan int64) {

	go func() {
		for _, x := range array {
			chIn <- x
		}
		close(chIn)
	}()
}

// MultBy2 получает на вход 2 канала chIn и chOut. В горутине читаем данные из канала chIn, производим метематическую операцию x*2 и результат операции пишем в канал chOut.
func MultBy2(chIn, chOut chan int64) {
	go func() {
		for x := range chIn {
			chOut <- x * 2
		}
		close(chOut)
	}()
}

// PrintResult читает данные из канала chOut и выводим в stdout.
func PrintResult(chOut chan int64) {
	for res := range chOut {
		fmt.Println(res)
	}
}

func main() {
	// chIn канал в который будут писатся числа из массива.
	chIn := make(chan int64)
	// chOut канал из которого будем производиться чтение.
	chOut := make(chan int64)

	// массив чисел.
	array := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	WriteCh(array, chIn)
	MultBy2(chIn, chOut)
	PrintResult(chOut)
}
