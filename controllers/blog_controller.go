package controllers

import (
	"fmt"
	"hello/models"
	"hello/utils"
	_ "strings"

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

// 跨域
func (c *BlogController) AllowCross() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")       //允许访问源
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")    //允许post访问
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization") //header的类型
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Max-Age", "1728000")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Ctx.ResponseWriter.Header().Set("content-type", "application/json") //返回数据格式是json
}

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

func (c *BlogController) ChangeOne() {
	ids := make([]string, 0, 13)
	c.Ctx.Input.Bind(&ids, "ids")
	fmt.Println(ids)
	blog_inst := models.Blog{}
	blogs := blog_inst.ChangeSome(ids, 1)
	utils.Logger.Debug("find %d blog", len(blogs))
	blog := blogs[0]

	blog_map := blog_inst.ChangeToMapOne(blog)

	c.Data["json"] = &blog_map
	c.ServeJSON()
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

// 根据id 查找对应blog，json版
func (c *BlogController) GetDetailById() {
	//	id := c.GetString("id")
	id := c.Ctx.Input.Param(":id")
	blog_inst := models.Blog{}
	blog := blog_inst.GetById(id)

	res := make(map[string]string)

	res["Title"] = blog.Title
	res["Content"] = blog.Content
	res["Id"] = blog.Id.Hex()

	c.Data["json"] = &res
	c.ServeJSON()
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
