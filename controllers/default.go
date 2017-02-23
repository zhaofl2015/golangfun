package controllers

import (
	"fmt"
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
	rotate_blog, wall_blog, window_blog := blog_inst.GetBlogForFirstPage()

	// 首页的最新blog的展示
	//	blog := blog_inst.GetNewestBlog()

	res_img_lst := GetFirstPageImages(7)

	// wall 展示

	wall_blog_map := blog_inst.ChangeToMapOne(wall_blog)
	wall_blog_map["img_url"] = res_img_lst[3]
	c.Data["wall_blog"] = wall_blog_map

	// 走马灯展示
	rotate_blog_map := blog_inst.ChangeToMap(rotate_blog)
	for ind, img_url := range res_img_lst[:3] {
		rotate_blog_map[ind]["img_url"] = img_url
		utils.Logger.Debug(img_url)
	}
	c.Data["rotate_blog"] = rotate_blog_map

	// 橱窗展示
	window_blog_map := blog_inst.ChangeToMap(window_blog)
	utils.Logger.Debug(fmt.Sprintf("redundant length : %d", len(res_img_lst[4:])))
	for ind, img_url := range res_img_lst[4:] {
		utils.Logger.Debug(fmt.Sprintf("ind %d value %s", ind, img_url))
		window_blog_map[ind]["img_url"] = img_url
	}
	c.Data["window_blog"] = window_blog_map

	c.Data["Author"] = beego.AppConfig.String("author")
	c.Data["Website"] = beego.AppConfig.String("website")
	c.Data["Email"] = beego.AppConfig.String("email")
	c.TplName = "welcome.html"

}

func GetFirstPageImages(count int) []string {
	// 找到对应的图片，用于展示
	rand_img_url := utils.GetRandomImageLocal(count)
	//	utils.Logger.Debug(strings.Join(rand_img_url, ","))
	return rand_img_url
}
