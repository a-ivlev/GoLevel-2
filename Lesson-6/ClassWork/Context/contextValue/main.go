// Демонстрация работы контекста вариант: Передача значений по контексту.

package main

import (
	"context"
	"fmt"
)

func main() {
	ctx0 := context.WithValue(context.Background(), "key", "value")
	ctx1 := context.WithValue(ctx0, "key", "value1")
	ctx2 := context.WithValue(ctx1, "Geek", "Brains")

	// Context штука слоистая, и мы не перезатираем слой, а добавляем слой.
	// И на предыдущих слоях остаётся всё попрежнему.
	fmt.Println("ctx0 key =", ctx0.Value("key")) // ctx0 key = value
	fmt.Println("ctx1 key =", ctx1.Value("key")) // ctx1 key = value1
	fmt.Println("ctx2 key =", ctx2.Value("key"), "ctx2 Geek =", ctx2.Value("Geek")) // ctx2 key = value1 ctx2 Geek = Brains
}
