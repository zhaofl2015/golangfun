package models

import (
	"log"

	"gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
)

type Blog struct {
	Title   string
	Content string
}

func (b Blog) Get_newest_blog() *Blog {
	session, error := mgo.Dial("10.200.8.127:27017")
	if error != nil {
		panic(error)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("newblog").C("blog")
	result := Blog{}
	error = c.Find(nil).Sort("-create_time").One(&result)
	if error != nil {
		log.Fatal(error)
	}

	return &result
}
