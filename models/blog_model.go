package models

import (
	"fmt"
	"log"
	"time"

	"hello/models"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BlogUser struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	Username   string        `json:"username" bson: "username"`
	Nickname   string        `json:"nickname" bson: "nickname"`
	Email      string        `json:"email" bson: "email"`
	IsAdmin    bool          `json:"is_admin" bson:"is_admin"`
	Privileges []int         `json:"privileges" bson: "privileges"`
	CreatedAt  time.Time     `json:"created_at" bson: "created_at"`
}

type Blog struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	Title      string        `json:"title" bson:"title"`
	Content    string        `json:"content" bson:"content"`
	Author     bson.ObjectId `json:"author" bson:"author"`
	Comment    []string      `json:"comment" bson:"comment"`
	CreateTime time.Time     `json:"create_time", bson:"create_time"`
	ViewCount  int           `json:"view_count" bson:"view_count"`
	Visible    int           `json:"visible" bson:"visible"`
	Tags       []string      `json:"tags" bson: "tags"`
	LastViewIP string        `json:"last_view_ip" bson:"last_view_ip"`
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

func (b Blog) Get_list(page int, per_page int) ([]Blog, int) {
	//	session, err := mgo.Dial(beego.AppConfig.String("mongoaddr"))
	//	if err != nil {
	//		panic(err)
	//	}

	//	defer session.Close()
	//	session.SetMode(mgo.Monotonic, true)
	c := session.DB(beego.AppConfig.String("mongodb")).C(beego.AppConfig.String("mongoblogcollection"))

	var result []Blog
	err = c.Find(bson.M{}).Sort("-_id").All(&result)
	total := len(result)
	left_side := per_page * (page - 1)
	if left_side < total {
		right_side := per_page * page
		if right_side > total {
			right_side = total
		}
		result = result[left_side:right_side]
	}
	//	fmt.Println(result)

	return result, total
}

func (b Blog) ChangeToMapOne(blog Blog) map[interface{}]interface{} {
	session, err := mgo.Dial(beego.AppConfig.String("mongoaddr"))
	if err != nil {
		panic(err)
	}

	defer session.Close()
	var res map[interface{}]interface{}

	return res
}

func (b Blog) ChangeToMap(blogs []Blog) []map[interface{}]interface{} {
	session, err := mgo.Dial(beego.AppConfig.String("mongoaddr"))
	if err != nil {
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(beego.AppConfig.String("mongodb")).C(beego.AppConfig.String("mongobloguser"))

	var res_list []map[interface{}]interface{}

	var user BlogUser

	for _, blog := range blogs {
		err = c.FindId(blog.Author).One(&user)
		//		user.CreatedAt = time.Parse(user.CreatedAt, "2006-01-02 15:04:05")
		fmt.Println(user)
		if err != nil {
			panic(err)
		}
		res := make(map[interface{}]interface{})
		res["Id"] = blog.Id.Hex()
		res["Title"] = blog.Title
		res["Content"] = blog.Content
		res["Comment"] = blog.Comment
		res["Tags"] = blog.Tags
		res["ViewCount"] = blog.ViewCount
		res["Author"] = user.Username
		res["Visible"] = blog.Visible
		res["LastViewIP"] = blog.LastViewIP
		res["CreateTime"] = blog.CreateTime
		//		res["CreateTime"] = blog.CreateTime.Format("2006-01-02 15:04:05")
		res_list = append(res_list, res)
	}
	return res_list
}
