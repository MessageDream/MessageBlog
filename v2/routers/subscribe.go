package routers

import (
	//"log"
	"regexp"
)

import (
	"../common"
	"../models"
)

type SubscribeRouter struct {
	baseRouter
}

type responsData struct {
	Status string
	ErrMsg string
}

func (this *SubscribeRouter) Get() {
	uid := this.Ctx.Input.Param(":uid")
	this.Data["json"] = uid
	this.ServeJson(true)
}

func (this *SubscribeRouter) Post() {
	email := this.GetString("email")
	rxOk := regexp.MustCompile(`\w[-._\w]*\w@\w[-._\w]*\w\.\w{2,3}`)

	if rxOk.MatchString(email) {
		subscription := models.Subscription{
			Email:  email,
			Uid:    common.UUID(),
			Status: false,
		}
		subscription.Set()
		common.SendMail("zmlv1314@163.com", "zm119www744dff15", "smtp.163.com:25", email, "MessageDream 博客订阅", "123456", "text")
		this.Data["json"] = responsData{
			Status: "ok",
		}
	} else {
		this.Data["json"] = responsData{
			ErrMsg: "邮箱地址格式不正确！",
		}
	}

	this.ServeJson(true)
}
