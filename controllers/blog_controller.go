package controllers

import (
	"hello/models"

	"github.com/astaxie/beego"
)

type BlogController struct {
	beego.Controller
}

//func (this *BlogController) Prepare() {
//	sess, _ := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
//	defer sess.SessionRelease(this.Ctx.ResponseWriter)
//	sess_username := sess.Get("username")
//	if sess_username == nil {
//		this.Ctx.Redirect(302, "/login")
//		return
//	}
//	this.Data["Username"] = sess_username
//}

// 获取最新的日志
func (c *BlogController) Get() {
	flash := beego.ReadFromRequest(&c.Controller)
	_, _ = flash.Data["notice"]
	blog_inst := models.Blog{}
	blog := blog_inst.GetNewestBlog()

	c.Data["Title"] = blog.Title
	c.Data["Content"] = blog.Content
	c.Data["Id"] = blog.Id
	c.TplName = "news.html"
}

// 使用json来交互信息
func (c *BlogController) GetNews() {
	flash := beego.ReadFromRequest(&c.Controller)
	_, _ = flash.Data["notice"]
	blog_inst := models.Blog{}
	blog := blog_inst.GetNewestBlog()

	res := make(map[string]string)

	res["Title"] = blog.Title
	res["Content"] = blog.Content
	res["Id"] = blog.Id.Hex()

	c.Data["json"] = &res
	c.ServeJSON()
}

func (c *BlogController) GetVue() {
	c.TplName = "blog_detail.html"
}

// 根据id查找对应的blog
func (c *BlogController) GetById() {
	id := c.GetString("id")
	blog_inst := models.Blog{}
	blog := blog_inst.GetById(id)
	c.Data["Title"] = blog.Title
	c.Data["Content"] = blog.Content
	c.Data["Id"] = blog.Id
	c.TplName = "news.html"
}

func (c *BlogController) List() {
	page, err := c.GetInt("page")
	if err != nil {
		page = 1
	}
	per_page := 10
	blog_inst := models.Blog{}
	//	fmt.Println(page, per_page)
	result, total := blog_inst.GetList(page, per_page)
	data := blog_inst.ChangeToMap(result)
	c.Data["data"] = data
	//	fmt.Println(data[0])
	c.Data["total"] = total
	c.Data["page"] = page
	c.Data["per_page"] = per_page
	c.TplName = "blog_list.html"
}
