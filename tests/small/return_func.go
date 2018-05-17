package main

import (
	"fmt"
)

func abc() (int, int) {
	return 1, 2
}

func main() {
	fmt.Println(abc()[0])
}
