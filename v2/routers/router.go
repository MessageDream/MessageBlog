package routers

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

type CurrentCategoryInfo struct {
	Title  string
	Name   string
	NName  string
	NTitle string
	ATitle string

	CurrentCat models.Category
	CurrentNod models.Node
}

type T_DATA struct {
	Flags T_FLAGS
	Vars  interface{}
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
	//	config := common.Webconfig
	data := T_DATA{
		Vars: vars,
	}
	return data
}

//获取当前网站的标题
func (c *CurrentCategoryInfo) GetCurrentTitle() string {
	if c.ATitle != "" {
		return c.ATitle
	} //文章标题
	cats := models.Categories
	if c.Name != "" {
		for _, v := range cats {
			if v.Name == c.Name {
				c.CurrentCat = v
				break
			}

		}
		return c.CurrentCat.Title
	} //分类标题

	if c.NName != "" {
		for _, v := range cats {
			flag := false
			for _, vn := range v.Nodes {
				if vn.Name == c.NName {
					flag = true
					c.CurrentNod = vn
					break
				}
			}
			if flag {
				c.CurrentCat = v
				break
			}
		}
		return c.CurrentNod.Title
	} //节点标题

	return ""
}

//获取分类名称
func (c *CurrentCategoryInfo) GetCName() string {
	if c.Name != "" {
		return c.Name
	}

	if &c.CurrentCat == nil {
		return c.CurrentCat.Name
	}
	cats := models.Categories
	if c.NName != "" {
		var name string = ""
		for _, v := range cats {
			flag := false
			for _, vn := range v.Nodes {
				if vn.Name == c.NName {
					flag = true
					break
				}
			}
			if flag {
				name = v.Name
				break
			}
		}
		return name
	}

	return ""
}

func (this *baseRouter) Prepare() {
	if this.Ctx.Request.Method == "GET" {
		this.Data["TagList"], _, _ = models.GetTags(&bson.M{}, 0, 10, "")
		this.Data["HotList"], _, _ = models.GetArticles(&bson.M{}, 0, 10, "-views")
		this.Data["RecentList"], _, _ = models.GetArticles(&bson.M{}, 0, 10, "-createdtime")
		this.Data["CategoryList"] = models.Categories
		this.Data["SiteConfig"] = common.Webconfig
	}
}

func (this *baseRouter) CheckError(err error) bool {
	if err != nil && err.Error() == "not found" {
		this.Redirect("/prompt/404", 302)
		return false
	} else if err != nil {
		beego.Error(err)
		return false
	}
	return true
}
