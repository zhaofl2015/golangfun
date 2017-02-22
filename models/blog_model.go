package models

import (
	"errors"
	"hello/utils"
	"time"

	"github.com/astaxie/beego"
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

var (
	BlogDBName         = beego.AppConfig.String("mongodb")
	BlogCollection     = beego.AppConfig.String("mongoblogcollection")
	BlogUserCollection = beego.AppConfig.String("mongobloguser")
)

// 获取最新的日志
func (b Blog) GetNewestBlog() *Blog {
	ds := NewDataStore()
	defer ds.Close()
	c := ds.C(BlogDBName, BlogCollection)
	result := Blog{}
	err := c.Find(nil).Sort("-create_time").One(&result)
	if err != nil {
		utils.Logger.Critical("cannot find blog")
	}
	return &result
}

//根据id获取日志
func (b Blog) GetById(id string) *Blog {
	ds := NewDataStore()
	defer ds.Close()
	c := ds.C(BlogDBName, BlogCollection)

	result := Blog{}
	c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
	return &result
}

// 批量获取日志
func (b Blog) GetList(page int, per_page int) ([]Blog, int) {
	utils.Logger.Debug("getting the list")

	ds := NewDataStore()
	defer ds.Close()
	c := ds.C(BlogDBName, BlogCollection)

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

// 将blog转换为map
func (b Blog) ChangeToMapOne(blog Blog) map[interface{}]interface{} {
	ds := NewDataStore()
	defer ds.Close()
	c := ds.C(BlogDBName, BlogUserCollection)

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
	res["PureContent"] = utils.RemoveHtmlTag(blog.Content)
	res["NewUrl"] = "/getone?id=" + blog.Id.Hex()

	preview_count, _ := beego.AppConfig.Int("first_page_preview_count")
	var pure_content []rune
	if content_string, ok := res["PureContent"].(string); ok {
		pure_content = []rune(content_string)
	} else {
		panic(errors.New("cannot transfer to string"))
	}

	if len(pure_content) < preview_count {
		preview_count = len(pure_content)
	}

	res["Summary"] = string(pure_content[:preview_count])

	return res
}

// 批量将blog转换为map格式
func (b Blog) ChangeToMap(blogs []Blog) []map[interface{}]interface{} {
	var res_list []map[interface{}]interface{}

	for _, blog := range blogs {
		res := b.ChangeToMapOne(blog)
		res_list = append(res_list, res)
	}
	return res_list
}

// 获取首页需要的blog：走马灯，大板块，橱窗
func (b Blog) GetBlogForFirstPage() (rotate []Blog, wall Blog, window []Blog) {
	//	found_ids := make(map[bson.ObjectId]int)
	ds := NewDataStore()
	defer ds.Close()
	c := ds.C(BlogDBName, BlogCollection)

	count := 7
	var blog_list []Blog
	iter := c.Find(nil).Limit(count).Iter()
	err := iter.All(&blog_list)
	if err != nil {
		utils.Logger.Error("can not find blog")
	}

	if len(blog_list) < count {
		panic(errors.New("no enough blogs"))
	}

	return blog_list[:3], blog_list[3], blog_list[4:]
}
