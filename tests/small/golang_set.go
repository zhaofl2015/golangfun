package main

import (
	"fmt"

	gset "github.com/deckarep/golang-set"
)

func main() {
	l := []int{1, 2, 3}
	n := []interface{}{}
	n = append(n, l...)
	//	for _, item := range l {
	//		n = append(n, item)
	//	}
	s := gset.NewSetFromSlice(n)
	//	var d []interface{}
	//	d, = l.()
	//	s := gset.NewSetFromSlice(l)
	fmt.Println(len(s.ToSlice()))
}
