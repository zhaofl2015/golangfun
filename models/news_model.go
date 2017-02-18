package models

import (
	"log"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Blog struct {
	Id      bson.ObjectId `json:"id" bson:"_id"`
	Title   string        `json:"title" bson:"title"`
	Content string        `json:"content" bson:"content"`
	Author  bson.ObjectId `json:"author" bson:"author"`
}

func (b Blog) Get_newest_blog() *Blog {
	session, error := mgo.Dial(beego.AppConfig.String("mongoaddr"))
	if error != nil {
		panic(error)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(beego.AppConfig.String("mongodb")).C(beego.AppConfig.String("mongoblogcollection"))
	result := Blog{}
	error = c.Find(nil).Sort("-create_time").One(&result)
	if error != nil {
		log.Fatal(error)
	}

	return &result
}

func (b Blog) Get_by_id(id string) *Blog {
	session, err := mgo.Dial(beego.AppConfig.String("mongoaddr"))
	if err != nil {
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(beego.AppConfig.String("mongodb")).C(beego.AppConfig.String("mongoblogcollection"))

	result := Blog{}
	err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
	return &result
}
