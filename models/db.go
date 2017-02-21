package models

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

//type DataStore struct {
//	session *mgo.Session
//}

//var ds DataStore

var (
	session *mgo.Session
)

func Session() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.Dial(beego.AppConfig.String("mongoaddr"))
		if err != nil {
			panic(err)
		}
		session.SetMode(mgo.Monotonic, true)
	}
	return session.Clone()
}

func Init() {
	//	ds = DataStore{}
	//	session, err := mgo.Dial(beego.AppConfig.String("mongoaddr"))
	//	if err != nil {
	//		panic(err)
	//	}
	//	ds.session = session
}

//func (s *DataStore) dataStore() *DataStore {
//	return &DataStore{s.session.Copy()}
//}
