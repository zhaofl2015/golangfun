package controllers

import (
	"fmt"
	"hello/models"

	"github.com/astaxie/beego"
)

type NewsController struct {
	beego.Controller
}

//func (this *NewsController) Prepare() {
//	sess, _ := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
//	defer sess.SessionRelease(this.Ctx.ResponseWriter)
//	sess_username := sess.Get("username")
//	if sess_username == nil {
//		this.Ctx.Redirect(302, "/login")
//		return
//	}
//	this.Data["Username"] = sess_username
//}

func (c *NewsController) Get() {
	blog_inst := models.Blog{}
	blog := blog_inst.Get_newest_blog()

	c.Data["Title"] = blog.Title
	c.Data["Content"] = blog.Content
	c.Data["Id"] = blog.Id
	//	c.Data["Title"] = "A light is on for ever"
	//	c.Data["Content"] = "As we know, a light is turned off for resource reasons, but today China's handsomest guy has a light is on for ever"
	c.TplName = "news.html"
}

func (c *NewsController) Get_by_id() {
	id := c.GetString("id")
	fmt.Println(id, "retrieve info")
	blog_inst := models.Blog{}
	blog := blog_inst.Get_by_id(id)
	c.Data["Title"] = blog.Title
	c.Data["Content"] = blog.Content
	c.Data["Id"] = blog.Id
	//	c.Data["Title"] = "A light is on for ever"
	//	c.Data["Content"] = "As we know, a light is turned off for resource reasons, but today China's handsomest guy has a light is on for ever"
	c.TplName = "news.html"
}
