package controllers

import (
	"hello/models"

	"github.com/astaxie/beego"
)

type NewsController struct {
	beego.Controller
}

func (c *NewsController) Get() {
	blog_inst := models.Blog{}
	blog := blog_inst.Get_newest_blog()

	c.Data["Title"] = blog.Title
	c.Data["Content"] = blog.Content
	//	c.Data["Title"] = "A light is on for ever"
	//	c.Data["Content"] = "As we know, a light is turned off for resource reasons, but today China's handsomest guy has a light is on for ever"
	c.TplName = "news.html"
}

func (c *NewsController) Get_by_id() {
	blog_inst := models.Blog{}
	blog := blog_inst.Get_by_id()
	c.Data["Title"] = blog.Title
	c.Data["Content"] = blog.Content
	//	c.Data["Title"] = "A light is on for ever"
	//	c.Data["Content"] = "As we know, a light is turned off for resource reasons, but today China's handsomest guy has a light is on for ever"
	c.TplName = "news.html"
}
