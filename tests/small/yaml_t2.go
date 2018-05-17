package main

import (
	"fmt"

	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func main() {
	data, err := ioutil.ReadFile("../question-bank/conf/config.yml")
	if err != nil {
		panic(err)
	}

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)
	fmt.Println(m["httpport"])
}
