package main

import (
	"fmt"
	//	"log"
	//	"os"

	"gopkg.in/olivere/elastic.v3"
)

var (
	elastic_client elastic.Client
)

type EsXxQuestionLogic struct {
	SubjectId     int
	ContentTypeId int
}

func main() {
	// TODO 修改信息为配置
	elastic_client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://192.168.100.14:9200"))
	if err != nil {
		panic(err)
	}

	info, code, err := elastic_client.Ping("http://192.168.100.14:9200").Do()
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s", code, info.Version.Number)

	es_q := EsXxQuestionLogic{SubjectId: 103, ContentTypeId: 1030232}

	put1, err := elastic_client.Index().Index("testq").Type("question").Id("1").BodyJson(es_q).Do()
	if err != nil {
		panic(err)
	}

	fmt.Println(put1)
}
