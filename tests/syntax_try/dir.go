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

<<<<<<< HEAD
	ch := make(chan int, 1)

	ch <- 1
	value := <-ch
	fmt.Println(value)
=======
	list := []int{1, 2, 3, 4, 5}
	for ind, val := range list[3:] {
		fmt.Println(ind, val)
	}

	fmt.Printf("%d", 4)
>>>>>>> 149b0a7b234a075583b0681a1f70031404e02fb5
}
