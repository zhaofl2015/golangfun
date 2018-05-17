package main

import "fmt"

func main() {
	a := make(map[int]int)

	a[2] += 1
	a[3] += 2
	fmt.Println(a)
	fmt.Println(bool(5 != 0))
	fmt.Println(!true)
}
