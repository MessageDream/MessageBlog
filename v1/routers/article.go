package routers

import (
	"../models"
)

import (
	"labix.org/v2/mgo/bson"
)

type ArticleRouter struct {
	baseRouter
}

func (this *ArticleRouter) Get() {
	name := this.Ctx.Input.Params(":article")
	article, err := models.GetArticle(&bson.M{"metadata.name": name})
	if !this.CheckError(err) {
		return
	}
	vars := make(map[string]interface{})
	vars["ArticleName"] = article.Metadata.Name
	vars["ArticleTitle"] = article.Metadata.Title
	prename, nextname, _ := models.GetAroundArticleName(&bson.M{"metadata.createdtime": &bson.M{"$lt": article.Metadata.CreatedTime}}, &bson.M{"metadata.createdtime": &bson.M{"$gt": article.Metadata.CreatedTime}})
	vars["PrevArticleName"] = prename
	vars["NextArticleName"] = nextname
	data := MakeData(vars)

	data.Flags.Single = true
	this.Data["Data"] = data
	this.Data["Article"] = article

	this.Layout = "layout.html"
	this.TplNames = "article.html"
}
