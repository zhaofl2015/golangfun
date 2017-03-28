package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
)

func main() {
	files, err := ioutil.ReadDir("../../static/img/firstpage")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	fmt.Println(rand.Perm(5))

	ch := make(chan int, 1)

	ch <- 1
	value := <-ch
	fmt.Println(value)
}
