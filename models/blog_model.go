package models

import (
	"time"

	"github.com/astaxie/beego"
	//	"gopkg.in/mgo.v2"
	"hello/utils"

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

//func (ds *DataStore) blog_collection() *mgo.Collection {
//	db := ds.session.DB(beego.AppConfig.String("mongodb"))
//	return db.C(beego.AppConfig.String("mongoblogcollection"))
//}

func (b Blog) GetNewestBlog() *Blog {
	session := Session()
	defer session.Close()
	c := session.DB(beego.AppConfig.String("mongodb")).C(beego.AppConfig.String("mongoblogcollection"))
	result := Blog{}
	err := c.Find(nil).Sort("-create_time").One(&result)
	if err != nil {
		utils.Logger.Critical("cannot find blog")
	}
	return &result
}

func (b Blog) GetById(id string) *Blog {
	session := Session()
	defer session.Close()
	c := session.DB(beego.AppConfig.String("mongodb")).C(beego.AppConfig.String("mongoblogcollection"))

	result := Blog{}
	c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
	return &result
}

func (b Blog) GetList(page int, per_page int) ([]Blog, int) {
	utils.Logger.Debug("getting the list")
	session := Session()
	defer session.Close()
	c := session.DB(beego.AppConfig.String("mongodb")).C(beego.AppConfig.String("mongoblogcollection"))
	if c == nil {
		utils.Logger.Error("collection find failed")
	}
	var result []Blog
	err := c.Find(bson.M{}).Sort("-_id").All(&result)
	if err != nil {
		panic(err)
	}
	total := len(result)
	left_side := per_page * (page - 1)
	if left_side < total {
		right_side := per_page * page
		if right_side > total {
			right_side = total
		}
		result = result[left_side:right_side]
	}

	return result, total
}

func (b Blog) ChangeToMapOne(blog Blog) map[interface{}]interface{} {
	session := Session()
	defer session.Close()
	c := session.DB(beego.AppConfig.String("mongodb")).C(beego.AppConfig.String("mongobloguser"))

	var res map[interface{}]interface{}

	var user BlogUser
	err := c.FindId(blog.Author).One(&user)
	if err != nil {
		utils.Logger.Warn("cannot find user")
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

	return res
}

func (b Blog) ChangeToMap(blogs []Blog) []map[interface{}]interface{} {
	var res_list []map[interface{}]interface{}

	var user BlogUser

	for _, blog := range blogs {
		res := b.ChangeToMapOne(blog)
		res_list = append(res_list, res)
	}
	return res_list
}
