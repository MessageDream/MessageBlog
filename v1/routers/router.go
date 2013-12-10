package routers

import (
	"log"
)

import (
	"../common"
	"../models"
)

import (
	"github.com/astaxie/beego"
	"labix.org/v2/mgo/bson"
)

type baseRouter struct {
	beego.Controller
}

type T_FLAGS struct {
	Home     bool
	Articles bool
	Single   bool
	Tag      bool
	Page     bool
	Feed     bool

	WriterOverview bool
	WriterPages    bool
	WriterTags     bool
	WriterComments bool
	WriterSettings bool
	WriterEditor   bool
}

type T_DATA struct {
	Flags      T_FLAGS
	SiteConfig common.Config
	Vars       interface{}
}

var (
	ArticleCount int
	TagCount     int
)

// func init() {
// 	ArticleCount = models.GetArticleCount()
// 	TagCount = models.GetTagCount()
// }

func MakeData(vars interface{}) T_DATA {
	config := common.Webconfig
	data := T_DATA{
		SiteConfig: *config,
		Vars:       vars,
	}
	return data
}

func (this *baseRouter) Prepare() {
	this.Data["TagList"], _, _ = models.GetTags(&bson.M{}, 0, 10, "")
}

func (this *baseRouter) CheckError(err error) bool {
	if err != nil && err.Error() == "not found" {
		this.Redirect("/prompt/404", 302)
		return false
	} else if err != nil {
		log.Println(err)
		return false
	}
	return true
}
