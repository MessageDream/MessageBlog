package routers

import (
	//"html/template"
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

type HomeRouter struct {
	baseRouter
}

func (this *HomeRouter) Get() {

	limit := common.Webconfig.PageCount
	page, err := strconv.Atoi(this.Ctx.Input.Param(":page"))
	if err != nil {
		page = 1
	}

	vars := make(map[string]interface{})

	vars["Offset"] = page

	article, total, err := models.GetArticles(&bson.M{}, (page-1)*limit, limit, "-createdtime")

	if !this.CheckError(err) {
		return
	}

	if (page-1)*limit > total {
		this.Redirect("/prompt/404", 302)
		return
	}

	totalpage := int(math.Ceil(float64(total) / float64(limit)))
	vars["Pager"] = common.GetPager("page", page, totalpage)
	vars["CurrentCategory"] = &CurrentCategoryInfo{
		Name: "home",
	}
	data := MakeData(vars)

	this.Data["Data"] = data
	this.Data["Articles"] = article
	// this.Data["json"], _ = models.GetAllCategory()
	// this.ServeJson(true)
	this.Layout = "layout.html"
	this.TplNames = "articles.html"

}
