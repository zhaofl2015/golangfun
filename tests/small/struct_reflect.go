package main

import (
	"fmt"
	"reflect"
)

type A struct {
}

func (a A) FuncB() int {
	fmt.Println("execute func B")
	return 8
}

func (a A) FuncC() int {
	fmt.Println("execute func C")
	return 5
}

func main() {
	fmt.Println("main")
	aType := reflect.TypeOf(A{})
	aValue := reflect.ValueOf(A{})
	for i := 0; i < aType.NumMethod(); i++ {
		method := aType.Method(i)
		fmt.Println(method.Name)
		result := aValue.Method(i).Call([]reflect.Value{})
		fmt.Println(result[0].Interface())
	}
}
