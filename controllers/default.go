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

func (c *MainController) GetTmplate() {
	blog_inst := models.Blog{}
	// 展示所有的首页内容，走马灯，wall，橱窗
	rotate_blog, wall_blog, window_blog := blog_inst.GetBlogForFirstPage()

	utils.Logger.Debug("runing at %d", beego.BConfig.Listen.HTTPPort)

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

// 跨域
func (c *MainController) AllowCross() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")       //允许访问源
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")    //允许post访问
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization") //header的类型
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Max-Age", "1728000")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Ctx.ResponseWriter.Header().Set("content-type", "application/json") //返回数据格式是json
}

func (c *MainController) Get() {
	c.AllowCross()
	blog_inst := models.Blog{}
	// 展示所有的首页内容，走马灯，wall，橱窗
	rotate_blog, wall_blog, window_blog := blog_inst.GetBlogForFirstPage()

	utils.Logger.Debug("runing at %d", beego.BConfig.Listen.HTTPPort)

	// 首页的最新blog的展示
	//	blog := blog_inst.GetNewestBlog()

	res_img_lst := GetFirstPageImages(7)

	// wall 展示

	wall_blog_map := blog_inst.ChangeToMapOne(wall_blog)
	wall_blog_map["img_url"] = res_img_lst[3]

	// 走马灯展示
	rotate_blog_map := blog_inst.ChangeToMap(rotate_blog)
	for ind, img_url := range res_img_lst[:3] {
		rotate_blog_map[ind]["img_url"] = img_url
		utils.Logger.Debug(img_url)
	}

	// 橱窗展示
	window_blog_map := blog_inst.ChangeToMap(window_blog)
	utils.Logger.Debug(fmt.Sprintf("redundant length : %d", len(res_img_lst[4:])))
	for ind, img_url := range res_img_lst[4:] {
		utils.Logger.Debug(fmt.Sprintf("ind %d value %s", ind, img_url))
		window_blog_map[ind]["img_url"] = img_url
	}

	res := make(map[string]interface{})
	res["rotate_blog"] = rotate_blog_map
	res["window_blog"] = window_blog_map
	res["wall_blog"] = wall_blog_map
	res["Author"] = beego.AppConfig.String("author")
	res["Website"] = beego.AppConfig.String("website")
	res["Email"] = beego.AppConfig.String("email")

	c.Data["json"] = &res
	c.ServeJSON()

}

func GetFirstPageImages(count int) []string {
	// 找到对应的图片，用于展示
	rand_img_url := utils.GetRandomImageLocal(count)
	//	utils.Logger.Debug(strings.Join(rand_img_url, ","))
	return rand_img_url
}
