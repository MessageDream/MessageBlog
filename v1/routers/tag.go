package routers

import (
	"math"
)

import (
	"../models"
)

import (
	"labix.org/v2/mgo/bson"
)

type TagRouter struct {
	baseRouter
}

func (this *TagRouter) Get() {
	limit := 2
	tagname := this.Ctx.Input.Params(":tag")
	page, _ := this.GetInt("pos")

	articles, total, err := models.GetArticlesByTag(&bson.M{"name": tagname}, int(page)*limit, limit, "")

	if !this.CheckError(err) {
		return
	}

	if int(page)*limit > total {
		this.Redirect("/prompt/404", 302)
		return
	}

	vars := make(map[string]interface{})
	vars["TagName"] = tagname

	if page == 0 {
		vars["AtBegin"] = true
	}
	if int64(math.Ceil(float64(total)/float64(limit))) == page+1 {
		vars["AtEnd"] = true
	}

	vars["PrevTLPos"] = page - 1
	vars["NextTLPos"] = page + 1
	data := MakeData(vars)

	data.Flags.Tag = true
	this.Data["Data"] = data
	this.Data["Articles"] = articles

	this.Layout = "layout.html"
	this.TplNames = "tag.html"
}
