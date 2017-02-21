package models

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/session"
)

var session *session.Session

func Init() {
	session, err = mgo.Dail(beego.AppConfig.String("mongoaddr"))
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
}
