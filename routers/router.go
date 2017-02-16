package routers

import (
	"hello/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/news", &controllers.NewsController{})
	beego.Router("/getone", &controllers.NewsController{}, "get:Get_by_id")
}
