/*
Написать функцию, которая принимает на вход имя файла и название функции. Необходимо
подсчитать в этой функции количество вызовов асинхронных функций. Результат работы
должен возвращать количество вызовов int и ошибку error. Разрешается использовать только
go/parser, go/ast и go/token.
*/

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {

	count, err := CountGorutines("./gorutine.go", "rout")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(count)

}

func CountGorutines(filePath string, funcName string) (int, error) {
	var count int
	fset := token.NewFileSet()
	// парсим файл, чтобы получить AST
	astFile, err := parser.ParseFile(fset, filePath, nil, 0)
	if err != nil {
		return 0, err
	}

	for _, decl := range astFile.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		if funcDecl.Name.String() == funcName {
			for _, elem := range funcDecl.Body.List {
				_, ok := elem.(*ast.GoStmt)
				if ok {
					count++
				}
			}

		}
	}

	return count, nil
}
