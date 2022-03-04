// Демонстрация работы контекста вариант deadlock!

package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	ctx0 := context.Background()//context.WithCancel(context.Background())
	ctx1, cancel1 := context.WithCancel(ctx0)
	ctx2 := context.WithValue(ctx1, "Geek", "Brains")

	// Создаём не по значению т.к. будет использоваться в замыкании.
	// А замыкание переменные захватывает по значению.
	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()

		<-ctx0.Done()
		fmt.Println("Context ctx0 done, value \"Geek\"", ctx0.Value("Geek"))
	}()

	go func() {
		defer wg.Done()

		<-ctx2.Done()
		fmt.Println("Context ctx2 done, value \"Geek\"", ctx2.Value("Geek"))
	}()

	// Чтобы всё отработало правильно нужно завершать ctx0, вызовом cancel0().
	cancel1()

	// Получаем deadlock! Потому, что у нас завершился ctx1 и ctx2, 
	// но не завершился ctx0 и он никогда не завершится.
	wg.Wait()
	// Контекст отменяется сверху вниз, если отменился верхний его дочерние тоже отменяются.
	// Если отменился дочерний верхний не отменяется.
}