package main

import (
	"fmt"
	"sync"
)

type WP interface {
	Do(string)
	Close()
}

type Task = string
type WorkerPool struct {
	taskCh  chan Task
	wg      *sync.WaitGroup
}

func NewWorkerPool(workers int) *WorkerPool {
	taskCh := make(chan Task, workers)

	wp := &WorkerPool{
		taskCh: taskCh,
		wg: &sync.WaitGroup{},
	}

	wp.wg.Add(workers)
	for i := 1; i <= workers; i++ {
		go func(i int) {
			defer wp.wg.Done()
			for task := range wp.taskCh {
				fmt.Printf("Worker %d performs a task %s\n", i, task)
			}
		}(i)
	}

	return wp
}

func (wp *WorkerPool) Do(f string) {

	wp.taskCh <- f
}

func (wp *WorkerPool) Close() {

	// Если будет wp.wg.Wait(), а потом close(wp.taskCh), то мы получим deadlock!
	// Так как wp.wg.Wait() никогда не выполниться.
	close(wp.taskCh)
	
	wp.wg.Wait()
}

// func printNum(i int) func() {
// 	return func() {
// 		fmt.Printf("I'm printing i = %d\n", i)
// 	}

// }

func printNum(i int) string {
	return fmt.Sprintf("%d", i)
}

func main() {
	var wp  WP = NewWorkerPool(8)
	defer wp.Close()

	for i := 1; i <= 100; i++ {
		wp.Do(printNum(i))
	}
}
