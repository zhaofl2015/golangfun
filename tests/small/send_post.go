package main

import (
	"fmt"
	//	"fmt"
	//	"question-bank/utils"
	"question-bank/middle"

	"github.com/astaxie/beego"
)

func main() {
	//	data := make(map[string]string)
	//	data["key"] = "key"
	//	data["value"] = "value"
	//	headers := make(map[string]string)

	//	headers["Content-Type"] = "application/x-www-form-urlencoded"
	//	_, body := utils.HttpPostForm("http://127.0.0.1:5000/service/test_post", data, headers)
	//	fmt.Println(body)
	fmt.Println(beego.AppConfig.String("mongo_question_addr"))
	middle.SendFlush([]string{}, []string{"Q_10300228821456-4"}, []string{},
		[]string{}, []string{}, []string{}, []string{}, []string{}, []string{})
	//	a := []string{"1"}
	//	fmt.Println(a[-1])
}
