/*
Для самостоятельного изучения. Исправить исходный код программы таким образом, чтобы
команда vet не ругалась.
*/

package main

import (
	"encoding/json"
	"fmt"
)

type MyStruct struct{}

func (f MyStruct) MarshalJSON() ([]byte, error) {
	return []byte(`{"a": 0}`), nil
}

func main() {
	j, _ := json.Marshal(MyStruct{})
	fmt.Println(string(j))
}
