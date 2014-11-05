package root

import (
	"../../common"
)

type LoginRouter struct {
	rootBaseRouter
}

func (this *LoginRouter) Get() {
	sess_username, _ := this.GetSession("username").(string)
	//如果未登录
	if sess_username == "" {
		username, _ := this.Ctx.Request.Cookie("username")
		if username != nil {
			this.Data["UserName"] = username.Value
		}
		this.TplNames = "root/login.html"
	} else { //如果已登录
		this.Ctx.Redirect(302, "/root")
	}

}

func (this *LoginRouter) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	rem := this.GetString("remember")
	if username != "" && password != "" {
		if username == common.Webconfig.ManagerInfo.UserName && password == common.Webconfig.ManagerInfo.Password {
			if rem == "remember" {
				this.Ctx.SetCookie("username", username, 0)
			} else {
				this.Ctx.SetCookie("username", username, 1)
			}
			//登录成功设置session
			this.SetSession("username", username)

			this.Ctx.Redirect(302, "/root")
		} else {

			this.Ctx.Redirect(302, "/root/login")
		}
	} else {

		this.Ctx.Redirect(302, "/root/login")
	}
}
