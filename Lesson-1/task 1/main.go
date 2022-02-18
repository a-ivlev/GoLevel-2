package main

import "fmt"

func implicitPanic() {

	var a int
	fmt.Println(1 / a)
}

func main() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("panic recovered:", v)
		}
	}()

	implicitPanic()
}
