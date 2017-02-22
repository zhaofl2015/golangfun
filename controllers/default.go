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
	// 展示所有的首页内容，走马灯，wall，橱窗
	//	rotate_blog, wall_blog, window_blog := blog_inst.GetBlogForFirstPage()
	_, wall_blog, _ := blog_inst.GetBlogForFirstPage()

	// 首页的最新blog的展示
	//	blog := blog_inst.GetNewestBlog()

	c.Data["Title"] = wall_blog.Title
	// 计算首页的预览的长度
	preview_count, _ := beego.AppConfig.Int("first_page_preview_count")
	if len([]rune(utils.RemoveHtmlTag(wall_blog.Content))) < preview_count {
		preview_count = len([]rune(utils.RemoveHtmlTag(wall_blog.Content)))
	}
	c.Data["Content"] = string([]rune(utils.RemoveHtmlTag(wall_blog.Content))[:preview_count])
	c.Data["Id"] = wall_blog.Id
	c.Data["NewUrl"] = "/getone?id=" + wall_blog.Id.Hex()

	c.Data["Author"] = beego.AppConfig.String("author")

	c.Data["Website"] = beego.AppConfig.String("website")
	c.Data["Email"] = beego.AppConfig.String("email")
	c.TplName = "welcome.html"

}
