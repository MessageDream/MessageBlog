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
	"labix.org/v2/mgo/bson"
)

type TagRouter struct {
	baseRouter
}

func (this *TagRouter) Get() {
	limit := common.Webconfig.PageCount
	tagname := this.Ctx.Input.Param(":tag")
	page, err := strconv.Atoi(this.Ctx.Input.Param(":page"))
	if err != nil {
		page = 1
	}
	articles, total, err := models.GetArticlesByTag(&bson.M{"name": tagname}, (page-1)*limit, limit, "")

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
		ATitle: tagname,
	}
	vars["Pager"] = common.GetPager("tag/"+tagname, page, totalpage)
	data := MakeData(vars)

	data.Flags.Tag = true
	this.Data["Data"] = data
	this.Data["Articles"] = articles

	this.Layout = "layout.html"
	this.TplNames = "articles.html"
}
