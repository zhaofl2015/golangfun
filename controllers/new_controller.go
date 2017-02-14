package controllers

import (
	"news/Blog"

	"github.com/astaxie/beego"
)

type NewsController struct {
	beego.Controller
}

func (c *NewsController) Get() {
	blog := Blog.get_newest_blog()

	c.Data["Title"] = blog.title
	c.Data["Content"] = blog.content
	//	c.Data["Title"] = "A light is on for ever"
	//	c.Data["Content"] = "As we know, a light is turned off for resource reasons, but today China's handsomest guy has a light is on for ever"
	c.TplName = "news.html"
}
