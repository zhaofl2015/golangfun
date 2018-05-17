package main

import (
	"fmt"
	"reflect"
)

func in_array(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}
func main() {
	fmt.Println(in_array("hahaha", []string{"hahaha"}))
	fmt.Println(in_array(0, []int{1, 2, 3, 4, 5, 0}))
}
