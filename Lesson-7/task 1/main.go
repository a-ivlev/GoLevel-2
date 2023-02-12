/*
Написать функцию, которая принимает на вход структуру in (struct или кастомную struct) и
values map[string]interface{} (key - название поля структуры, которому нужно присвоить value
этой мапы). Необходимо по значениям из мапы изменить входящую структуру in с помощью
пакета reflect. Функция может возвращать только ошибку error. Написать к данной функции
тесты (чем больше, тем лучше - зачтется в плюс).
*/

package main

import (
	"errors"
	"fmt"
	"reflect"
)

type rect struct {
	weight int
	height int
}

type circle struct {
	radius int
}

func Set(v interface{}, src map[string]interface{}) error {
	rvp := reflect.ValueOf(v)

	if rvp.Kind() != reflect.Ptr {
		return errors.New("Ожидается указатель на структуру.")
	}

	rv := rvp.Elem()
	i := rv.Interface()
	switch i.(type) {
	case rect:
		val := rect{src["weight"].(int), src["height"].(int)}
		rv.Set(reflect.ValueOf(val))
	case circle:
		val := circle{200}
		rv.Set(reflect.ValueOf(val))
	}

	return nil
}

func main() {
	src := map[string]interface{}{"radius": 100, "weight": 200, "height": 100}
	circ := circle{}
	rectagle := rect{}

	if err := Set(circ, src); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", circ)
	if err := Set(&rectagle, src); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", rectagle)

}
