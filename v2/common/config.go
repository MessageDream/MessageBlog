package common

import (
	"github.com/astaxie/beego"
)

type User struct {
	UserName string
	Password string
	EmailServer string
	EmailSender string
	EmailPwd string
}

type UploadConf struct {
	QiniuAccessKey string
	QiniuSecretKey string
}

type Config struct {
	// sys config
	Port        int
	Certificate string
	SiteBase    string
	SiteURL     string
	LogFile     string
	// site config
	SiteTitle    string
	Dbconn       string
	StaticURL    string
	PageCount    int
	ThemeName    string
	ManagerInfo  User
	UploadConfig UploadConf
}

var Webconfig *Config = nil

func init() {
	port, _ := beego.AppConfig.Int("httpport")
	pageCount, _ := beego.AppConfig.Int("pagecount")
	Webconfig = &Config{
		Port:      port,
		PageCount: pageCount,
		SiteBase:  beego.AppConfig.String("sitebase"),
		SiteURL:   beego.AppConfig.String("siteurl"),
		StaticURL: beego.AppConfig.String("staticurl"),
		LogFile:   beego.AppConfig.String("logfile"),
		SiteTitle: beego.AppConfig.String("appname"),
		Dbconn:    beego.AppConfig.String("dbconn"),
		ThemeName: beego.AppConfig.String("themename"),
		ManagerInfo: User{
			UserName: beego.AppConfig.String("username"),
			Password: beego.AppConfig.String("password"),
			EmailServer: beego.AppConfig.String("emailserver"),
			EmailSender: beego.AppConfig.String("emailsender"),
			EmailPwd : beego.AppConfig.String("emailpwd"),
		},
		UploadConfig: UploadConf{
			QiniuAccessKey: beego.AppConfig.String("qiniuaccesskey"),
			QiniuSecretKey: beego.AppConfig.String("qiniusecretkey"),
		},
	}

}
