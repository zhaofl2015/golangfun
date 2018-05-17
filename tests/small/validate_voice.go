package main

import (
	"fmt"
	"question-bank/utils"
	//	"runtime"
)

func main() {
	fmt.Println(utils.EngineValidate(`{"DisplayText": "Jsgf Grammar Template", "Grammar": "#JSGF V1.0 utf-8 cn;\ngrammar main;\npublic <sample>=\"<s>\"(Good afternoon Miss Li)\"</s>\";", "GrammarWeight": "", "Version": 1}`))
	//	runtime.Caller()
}
