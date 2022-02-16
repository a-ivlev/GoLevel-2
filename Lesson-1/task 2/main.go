package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrCalculation = errors.New("calculation error")

func implicitPanic(a int) (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("%s panic recovered: %s\n", time.Now().Format("[15:04:05 02-01-2006]"), p)
		}
	}()

	fmt.Println(10 / a)
	return
}

func main() {
	_ = implicitPanic(5)

	err := implicitPanic(0)
	if err != nil {
		fmt.Println(err)
	}
}
