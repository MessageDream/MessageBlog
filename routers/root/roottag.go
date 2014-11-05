package root

import (
	"../../models"
)

type RootTagRouter struct {
	rootBaseRouter
}

// func (this *RootTagRouter) Get() {
// 	this.Data["Category"] = models.Categories
// 	this.Data["currentitem"] = "category"
// 	this.Layout = "root/layout.html"
// 	this.TplNames = "root/category.html"
// }

func (this *RootTagRouter) Post() {
	title := this.GetString("title")
	if title == "" {
		this.Data["json"] = 0
	}
	cat := models.TagWrapper{
		Name:  title,
		Title: title,
	}
	err := cat.SetTag()
	if err != nil {

	}
	this.Data["json"] = 1
	this.ServeJson(true)
}
