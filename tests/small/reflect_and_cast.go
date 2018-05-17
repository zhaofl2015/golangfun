package main

import (
	"fmt"
	"reflect"
)

type TAT struct {
	Name string
	Age  int
}

func main() {
	t := TAT{
		"haha",
		15}
	var d interface{}
	d = t
	fmt.Println(d)
	fmt.Println(reflect.TypeOf(d))
	fmt.Println(reflect.ValueOf(d).FieldByName("Name").Interface())
	//	fmt.Println(reflect.ValueOf(d).Elem().FieldByName("Name").Interface())

	var f float64
	f = 1.5

}
