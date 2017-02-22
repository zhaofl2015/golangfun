package models

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

type DataStore struct {
	session *mgo.Session
}

//var ds DataStore
var (
	session *mgo.Session
)

func init() {
	var err error
	session, err = mgo.Dial(beego.AppConfig.String("mongoaddr"))
	if err != nil {
		panic(err)
	}
}

// Close mgo.Session
func (d *DataStore) Close() {
	d.session.Close()
}

// create a new datastore for each http request
func NewDataStore() *DataStore {
	ds := &DataStore{
		session: session.Copy(),
	}
	return ds
}

// get collection with dbname and collection name
func (d *DataStore) C(dbname string, colname string) *mgo.Collection {
	return d.session.DB(dbname).C(colname)
}
