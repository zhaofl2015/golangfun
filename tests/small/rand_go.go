package main

import (
	"fmt"
	"math/rand"
	"question-bank/utils"
)

func main() {
	fmt.Printf("%s_%d%0.8d%0.3d\n", "OL", 103, 1234122, rand.Intn(1000))
	fmt.Println(utils.ExtractSubjectId("Q_103123413214"))
}
