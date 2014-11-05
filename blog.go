package main

import (
	"./common"
	"./models"
	"./routers"
	"./routers/root"
)

import (
	"github.com/astaxie/beego"
)

func main() {
	models.InitDb()
	beego.SessionOn = true
	beego.BeeLogger.SetLogger("file", `{"filename":"`+common.Webconfig.LogFile+`"}`)
	beego.Info("begin setting router")

	beego.Router("/", &routers.HomeRouter{})
	beego.Router(`/page/:page([\d]+)`, &routers.HomeRouter{})

	beego.Router(`/article/:article([\w-]+)`, &routers.ArticleRouter{})

	beego.Router(`/tag/:tag([\S]+)/:page([\d]+)`, &routers.TagRouter{})
	beego.Router(`/tag/:tag([\S]+)`, &routers.TagRouter{})

	beego.Router(`/node/:node([\S]+)/:page([\d]+)`, &routers.NodeRouter{})
	beego.Router(`/node/:node([\S]+)`, &routers.NodeRouter{})

	beego.Router(`/category/:category([\S]+)/:page([\d]+)`, &routers.CategoryRouter{})
	beego.Router(`/category/:category([\S]+)`, &routers.CategoryRouter{})

	beego.Router("/subscribe", &routers.SubscribeRouter{})
	beego.Router(`/subscribe/:uid([\S]+)`, &routers.SubscribeRouter{})
	beego.Router(`/desubscribe/:uid([\S]+)`, &routers.SubscribeRouter{})

	beego.Router("/root", &root.RootMainRouter{})

	beego.Router("/root/category", &root.RootCategoryRouter{})

	beego.Router("/root/node", &root.RootNodeRouter{})

	beego.Router("/root/tag", &root.RootTagRouter{})

	beego.Router(`/root/article`, &root.RootArticleRouter{})
	//beego.Router(`/root/article/:opt([a-z]+)`, &root.RootArticleRouter{})
	beego.Router(`/root/article/:page([\d]+)`, &root.RootArticleRouter{})

	beego.Router("/root/upload", &root.RootUploadRouter{})

	beego.Router("/root/login", &root.LoginRouter{})
	beego.Router("/root/loginout", &root.LoginOutRouter{})

	beego.Router(`/prompt/:code([\w]+)`, &routers.PromptRouter{})

	beego.Info("starting server")
	beego.Run()
}
