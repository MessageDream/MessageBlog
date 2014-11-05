package routers

import (
	//"log"
	//"fmt"
	"html/template"
	//"os"
	"regexp"
	"strings"
)

import (
	"../common"
	"../models"
)

import (
	"github.com/astaxie/beego"
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
	var t *template.Template
	var err error
	sub := models.Subscription{
		Uid: uid,
	}
	if strings.Contains(this.Ctx.Request.URL.Path, "desubscribe") {
		sub.Status = false
		sub.UpdateState()
		t, err = template.New("name").Parse(`您已成功取消订阅，北飘漂博客新的动态将不会发送到您的邮箱。`)
	} else {
		sub.Status = true
		sub.UpdateState()
		t, err = template.New("name").Parse(`恭喜，您已成功订阅北飘漂博客，博客的最新动态将会及时发送到您的邮箱。`)
	}

	if err != nil {
		beego.Error(err)
		return
	}
	err = t.Execute(this.Ctx.ResponseWriter, nil)
	if err != nil {
		beego.Error(err)
		return
	}
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
		content := common.PaseHtml(common.SubConfirmContent, common.SubscribeEmail{
			StaticUrl:  common.Webconfig.StaticURL,
			WebUrl:     common.Webconfig.SiteURL,
			Uid:        subscription.Uid,
			MailSender: common.Webconfig.ManagerInfo.EmailSender,
		})
		common.SendMail(common.Webconfig.ManagerInfo.EmailSender, common.Webconfig.ManagerInfo.EmailPwd, common.Webconfig.ManagerInfo.EmailServer, email, "北飘漂订阅确认", *content, "html")
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
