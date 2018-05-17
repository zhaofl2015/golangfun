package main 

import (
	"gopkg.in/mgo.v2"
	"fmt"
)
import "gopkg.in/mgo.v2/bson"
import "question-bank/models"



func main(){
	ds := models.NewDataStore()
	c := ds.C("vox-question", "xx_questions")
	for doc := range c.Find(bson.M{"subject_id": 103}) {
		fmt.Println(doc.)
	}
}