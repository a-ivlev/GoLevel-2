/*
Задание 3
Смоделировать ситуацию “гонки”, и проверить программу на наличии “гонки”.

Проверить программу на наличие в коде состояния гонки следующей командой:
go run -race main.go
*/

package main

import (
	"fmt"
	"sync"
)

const count = 10

func main() {
	var (
		counter int
		wg      sync.WaitGroup
	)
	wg.Add(count)
	for i := 0; i < count; i += 1 {
		go func() {
			defer wg.Done()
			counter += 1
		}()
	}

	wg.Add(count)
	for i := 0; i < count; i += 1 {
		go func() {
			defer wg.Done()
			counter -= 1
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
