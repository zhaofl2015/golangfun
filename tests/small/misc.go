package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abc-cde"
	fmt.Println(s[:strings.Index(s, "-")])
}
