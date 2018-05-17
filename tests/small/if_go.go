package main

import (
	"fmt"
)
 
func A() (bool, string){
	return true, ""
}

func main() {
	if 5 == 5 && (b, c := A(); b==true) {
		fmt.Println("yes")
	}
}