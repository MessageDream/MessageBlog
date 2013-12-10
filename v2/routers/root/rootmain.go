package root

import (
	"../../models"
	"runtime"
)

type RootMainRouter struct {
	rootBaseRouter
}

func (this *RootMainRouter) Get() {
	dbinfo, _ := models.DB.Session.BuildInfo()
	this.Data["dbinfo"] = dbinfo.Version
	this.Data["remoteproto"] = this.Ctx.Request.Proto
	this.Data["remotehost"] = this.Ctx.Request.Host
	this.Data["remoteos"] = runtime.GOOS
	this.Data["remotearch"] = runtime.GOARCH
	this.Data["remotecpus"] = runtime.NumCPU()
	this.Data["golangver"] = runtime.Version()
	this.Data["currentitem"] = "index"
	this.Layout = "root/layout.html"
	this.TplNames = "root/index.html"
}
