package main

import (
	"os"
	"sync"
	"runtime/trace"
)

type Task = int64

type Result struct {
	worker int
	num    int64
	res    uint64
}

type WorkerPool struct {
	taskCh   chan Task
	resultCh chan Result
	wg       *sync.WaitGroup
	store    map[int64]Result
}

func NewWorkerPool(workers int, taskCh chan Task) *WorkerPool {
	//taskCh := make(chan Task, workers)
	resultCh := make(chan Result, workers)
	wg := &sync.WaitGroup{}
	wp := &WorkerPool{
		taskCh:   taskCh,
		resultCh: resultCh,
		wg:       wg,
		store:    make(map[int64]Result),
	}

	wp.wg.Add(workers)
	for i := 1; i <= workers; i++ {
		go func(i int) {
			defer wp.wg.Done()
			for task := range wp.taskCh {
				res := Result{
					worker: i,
					num:    task,
					res:    uint64(task * task),
				}
				wp.resultCh <- res
			}
		}(i)
	}

	go func() {
		wp.wg.Wait()
		close(wp.resultCh)
	}()

	return wp
}

func (wp *WorkerPool) DoTask(count int) {
	defer wp.wg.Done()
	defer close(wp.taskCh)
	for i := 1; i <= count; i++ {
		wp.taskCh <- int64(i)
	}
}

func (wp *WorkerPool) StoreResult() {

	for res := range wp.resultCh {
		//fmt.Printf("worker %d %d * %d = %d\n", res.worker, res.num, res.num, res.res)
		wp.store[res.num] = res
	}
}

func main() {
	// Функция Start начинает трассировку фрагмента исходного кода до тех пор,
	// пока не будет вызвана функция Stop. Трассировка пишется в экземпляр io.Writer,
	// переданный функции Start параметром. В данном случае мы пишем трассировку
	// в стандартный поток ошибок os.Stderr. Можно также ограничить число процессоров,
	// которые будут использованны при старте программы командой GOMAXPROCS=1 go run main.go 2>trace.out .
	// Затем файл может быть открыт при помощи инструмента trace: go tool trace trace.out
	trace.Start(os.Stderr)
	defer trace.Stop()
	count := 1000
	taskCh := make(chan Task, count)
	wp := NewWorkerPool(7, taskCh)

	wp.wg.Add(1)
	go wp.DoTask(count)

	wp.StoreResult()
}
