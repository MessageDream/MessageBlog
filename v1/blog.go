package main

import (
	"./models"
	"./routers"
)

import (
	"github.com/astaxie/beego"
)

func main() {
	models.InitDb("10.1.5.36:27017")
	//Init()
	//models.InitDb("localhost:27017")
	beego.HttpPort = 8888
	beego.SetStaticPath("/static", "./static")
	//beego.ViewsPath = "/views/" + common.Webconfig.ThemeName
	beego.Router("/", &routers.HomeRouter{})
	beego.Router("/?pos=:id int", &routers.HomeRouter{})
	beego.Router(`/:article([\w]+)`, &routers.ArticleRouter{})
	beego.Router(`/tag/:tag([\S]+)`+"?pos=:id int", &routers.TagRouter{})
	beego.Router(`/tag/:tag([^\?]+)`, &routers.TagRouter{})
	beego.Router(`/prompt/:code([\w]+)`, &routers.PromptRouter{})
	beego.Run()
}
