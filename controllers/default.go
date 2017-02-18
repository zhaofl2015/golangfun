package controllers

import (
	//	"fmt"
	"hello/models"
	"hello/utils"

	"github.com/astaxie/beego"
	//	"gopkg.in/mgo.v2/bson"
)

type MainController struct {
	beego.Controller
}

// 暂时关闭登录功能

//func (this *MainController) Prepare() {
//	sess, _ := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
//	defer sess.SessionRelease(this.Ctx.ResponseWriter)
//	sess_username := sess.Get("username")
//	if sess_username == nil {
//		this.Ctx.Redirect(302, "/login")
//		return
//	}
//	this.Data["Username"] = sess_username
//}

func (c *MainController) Get() {
	blog_inst := models.Blog{}
	blog := blog_inst.Get_newest_blog()

	c.Data["Title"] = blog.Title
	c.Data["Content"] = utils.RemoveHtmlTag(blog.Content)[:40]
	c.Data["Id"] = blog.Id
	c.Data["NewUrl"] = "/getone?id=" + blog.Id.Hex()
	//	fmt.Println(c.Data["NewUrl"])

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "welcome.html"
}
