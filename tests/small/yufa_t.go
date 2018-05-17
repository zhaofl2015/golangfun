package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	fmt.Println(reflect.TypeOf(1).Name())
	fmt.Println("==========")

	a := "CBC"

	switch {
	case strings.HasPrefix(a, "A"):
		fmt.Println("prefix A")
	case strings.HasPrefix(a, "B"):
		fmt.Println("prefix B")
	default:
		fmt.Println("not found")
	}

	b := []string{"aa", "bb", "cc"}
	step := 1
	i := 0
	for i < len(b) {
		fmt.Println(b[i : i+1])
		i += step
	}
	//	fmt.Println(b[1])
	fmt.Println(math.Abs(-1))

	var c []interface{}
	c = append(c, "aa")
	c = append(c, "bb")
	var d []string
	for _, item := range c {
		d = append(d, item.(string))
	}
	fmt.Println(d)
	var e []string
	fmt.Println(len(strings.Join(e, ",")))
}
