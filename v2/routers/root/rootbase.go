package root

import (
	"../../common"
	//"../../models"
)

import (
	"github.com/astaxie/beego"
	// "labix.org/v2/mgo/bson"
)

type rootBaseRouter struct {
	beego.Controller
}

func (this *rootBaseRouter) Prepare() {
	if this.Ctx.Request.Method == "GET" {
		this.Data["SiteConfig"] = common.Webconfig
	}
	if this.Ctx.Request.URL.Path != "/root/login" {
		sess_username, _ := this.GetSession("username").(string)
		this.Data["UserName"] = sess_username
		//如果未登录
		if sess_username == "" {
			this.Ctx.Redirect(302, "/root/login")
		}
	}

}
