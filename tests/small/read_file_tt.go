package main

import (
	"fmt"
	"question-bank/models"
)

func aaa() {
	doc_instance := models.XxQuestion{}
	doc := doc_instance.GetById("Q_10300403190462")
	fmt.Println(doc)
	fmt.Println(doc.Content.SubContents[0].SubContentTypeId)
}
