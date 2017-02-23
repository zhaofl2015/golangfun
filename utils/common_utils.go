package utils

import (
	"github.com/astaxie/beego"
)

var (
	RootPath       = beego.AppConfig.String("root_path")
	FirstImagePath = beego.AppConfig.String("static_img_first")
)
