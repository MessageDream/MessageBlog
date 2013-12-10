package routers

import (
	//"html/template"
	"math"
	//"strconv"
)

import (
	"../models"
)
import (
	"labix.org/v2/mgo/bson"
)

type HomeRouter struct {
	baseRouter
}

func (this *HomeRouter) Get() {

	// article1 := models.Article{
	// 	Metadata: models.ArticleMetadata{
	// 		Name:           strconv.Itoa(int(bson.Now().UnixNano())),
	// 		Author:         "Jayden",
	// 		IsPage:         false,
	// 		Title:          "爱情价更高",
	// 		Tags:           []string{"生命", "爱情"},
	// 		FeaturedPicURL: "",
	// 		Summary:        "好的",
	// 		CreatedTime:    bson.Now(),
	// 		ModifiedTime:   bson.Now(),
	// 		Hits:           0,
	// 	},
	// 	Text: template.HTML("生命诚可贵，爱情价更高！"),
	// }
	// models.CreatArticle(&article1)

	limit := 10
	page, err := this.GetInt("pos")
	if err != nil {
		page = 0
	}

	vars := make(map[string]interface{})

	vars["Offset"] = int(page)

	article, total, err := models.GetArticles(&bson.M{}, int(page)*limit, limit, "")

	if !this.CheckError(err) {
		return
	}

	if int(page)*limit > total {
		this.Redirect("/prompt/404", 302)
		return
	}

	if page == 0 {
		vars["AtBegin"] = true
	}
	if int64(math.Ceil(float64(total)/float64(limit))) == page+1 {
		vars["AtEnd"] = true
	}

	vars["PrevTLPos"] = page - 1
	vars["NextTLPos"] = page + 1
	data := MakeData(vars)

	data.Flags.Home = true

	this.Data["Data"] = data
	this.Data["Articles"] = article

	this.Layout = "layout.html"
	this.TplNames = "articles.html"

}
