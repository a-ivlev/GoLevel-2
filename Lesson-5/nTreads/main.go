// nTreads программа, которая запускает 𝑛 потоков и дожидается завершения их всех.
package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			fmt.Printf("Tread %d\n", j)
		}(i)
	}
	wg.Wait()
	fmt.Println("Done")
}
