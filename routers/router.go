package routers

import (
	"hello/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/hello", &controllers.MainController{}, "get:Get")

	beego.Router("/news", &controllers.BlogController{}, "get:Get")
	beego.Router("/news-2", &controllers.BlogController{}, "get:GetNews")
	beego.Router("/vue", &controllers.BlogController{}, "get:GetVue")
	beego.Router("/change-one", &controllers.BlogController{}, "post:ChangeOne")
	beego.Router("/getone", &controllers.BlogController{}, "get:GetById")
	beego.Router("/getdetail/:id", &controllers.BlogController{}, "get:GetDetailById")
	beego.Router("/blog-list", &controllers.BlogController{}, "get:List")
	beego.Router("/blog-list-api", &controllers.BlogController{}, "get:ListApi")

	beego.Router("/blog-months", &controllers.BlogController{}, "get:BlogMonths")

	beego.Router("/blog-tags", &controllers.BlogController{}, "get:BlogTagClouds")

	beego.Router("/login", &controllers.UserController{}, "get:Login;post:Post")
	beego.Router("/logout", &controllers.UserController{}, "get:Logout")
	beego.Router("/fun", &controllers.UserController{}, "get:Fun")
}
