package main

import (
	"fmt"
	"question-bank/utils"
)

func main() {
	a := []int{1, 2}
	b := []int{2, 1}
	c := []interface{}{3, 1}
	e := []interface{}{3, 1, 5, 2, 9, 11}
	d := []interface{}{1, 2, 3, 5, 9, 11}

	fmt.Println(utils.IsTwoSliceSame(utils.TransferArrayT2Interface(a), utils.TransferArrayT2Interface(b)))
	fmt.Println(utils.IsTwoSliceSame(c, utils.TransferArrayT2Interface(b)))
	fmt.Println(utils.IsTwoSliceSame(e, d))

	big := [][]int{[]int{0}, []int{1, 2}}
	var big_int []interface{}

	for _, l := range big {
		var item []interface{}
		for _, i := range l {
			item = append(item, i)
		}
		big_int = append(big_int, item)
	}

	fmt.Println("In Slice of slice", utils.IsSliceInSlice(utils.TransferArrayT2Interface(a), big_int))

	//	i := []int{1, 2, 3}
	//	fmt.Println([]int(i))
}
