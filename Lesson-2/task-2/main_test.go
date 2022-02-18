package main_test

func Example() {
	// chIn канал в который будут писатся числа из массива.
	chIn := make(chan int64)
	// chOut канал из которого будем производиться чтение.
	chOut := make(chan int64)

	// массив чисел.
	array := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}


	WriteCh(array, chIn)
	MultBy2(chIn, chOut)
	PrintResult(chOut)
	// Output:
	// 2
	// 4
	// 6
	// 8
	// 10
	// 12
	// 14
	// 16
	// 18
	// 20
	// 22
	// 24
}

func ExampleWriteCh(array []int64, chIn chan int64) {
	go func() {
		for _, x := range array {
			chIn <- x
		}
		close(chIn)
	}()
}

func ExampleMultBy2(chIn, chOut chan int64) {
	go func() {
		for x := range chIn {
			chOut <- x * 2
		}
		close(chOut)
	}()
}

func ExamplePrintResult(chOut chan int64) {
	for res := range chOut {
		fmt.Println(res)
	}
	// Output:
	// 2
	// 4
	// 6
	// 8
	// 10
	// 12
	// 14
	// 16
	// 18
	// 20
	// 22
	// 24
}
