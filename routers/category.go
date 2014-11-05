package routers

import (
	"math"
	"strconv"
)

import (
	"../common"
	"../models"
)

import (
	"gopkg.in/mgo.v2/bson"
)

type CategoryRouter struct {
	baseRouter
}

func (this *CategoryRouter) Get() {
	limit := common.Webconfig.PageCount
	categoryname := this.Ctx.Input.Param(":category")
	page, err := strconv.Atoi(this.Ctx.Input.Param(":page"))
	if err != nil {
		page = 1
	}

	articles, total, err := models.GetArticlesByNode(&bson.M{"cname": categoryname}, (page-1)*limit, limit, "-createdtime")

	if !this.CheckError(err) {
		return
	}

	if (page-1)*limit > total {
		this.Redirect("/prompt/404", 302)
		return
	}

	vars := make(map[string]interface{})
	totalpage := int(math.Ceil(float64(total) / float64(limit)))

	vars["CurrentCategory"] = &CurrentCategoryInfo{
		Name: categoryname,
	}
	vars["Pager"] = common.GetPager("category/"+categoryname, page, totalpage)
	data := MakeData(vars)

	this.Data["Data"] = data
	this.Data["Articles"] = articles
	// this.Data["json"] = articles
	// this.ServeJson(true)
	this.Layout = "layout.html"
	this.TplNames = "articles.html"
}
