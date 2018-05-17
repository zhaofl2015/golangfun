package main

import (
	"fmt"

	"github.com/satori/go.uuid"
)

func main() {
	fmt.Println(fmt.Sprintf("urn:uuid:%v", uuid.NewV1().String()))
	fmt.Println("abc" + "123")
}
