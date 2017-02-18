package routers

import (
	"hello/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get")
	beego.Router("/news", &controllers.NewsController{}, "get:Get")
	beego.Router("/getone", &controllers.NewsController{}, "get:Get_by_id")
	beego.Router("/login", &controllers.UserController{}, "get:Login;post:Post")
	beego.Router("/logout", &controllers.UserController{}, "get:Logout")
	beego.Router("/fun", &controllers.UserController{}, "get:Fun")
}
