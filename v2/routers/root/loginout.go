package root

type LoginOutRouter struct {
	rootBaseRouter
}

func (this *LoginOutRouter) Get() {
	this.DelSession("username")
	username, _ := this.Ctx.Request.Cookie("username")
	if username != nil {
		this.Data["UserName"] = username.Value
	}
	this.TplNames = "root/login.html"
}
