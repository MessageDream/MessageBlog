package root

import (
	"html/template"
	"strconv"
)

import (
	"../../common"
	"../../models"
)

import (
	"labix.org/v2/mgo/bson"
)

type RootArticleRouter struct {
	rootBaseRouter
}

func Publish(article *models.Article) {
	edata := common.SubscribeEmail{
		StaticUrl:   common.Webconfig.StaticURL,
		WebUrl:      common.Webconfig.SiteURL,
		MailSender:  common.Webconfig.ManagerInfo.EmailSender,
		Title:       (*article).Title,
		ArticleName: (*article).Name,
		Summary:     (*article).Summary,
	}
	var limit int
	for {
		offset := limit * 100
		subs, total, _ := models.GetSubscribes(&bson.M{"status": true}, offset, 100, "")
		for _, v := range *subs {
			edata.Uid = v.Uid
			content := common.PaseHtml(common.SubscribeEmailContent, edata)
			common.SendMail(common.Webconfig.ManagerInfo.EmailSender, common.Webconfig.ManagerInfo.EmailPwd, common.Webconfig.ManagerInfo.EmailServer, v.Email, "北飘漂博客发表了新文章", *content, "html")
		}
		if offset+100 >= total {
			break
		}
		limit++
	}
}

func (this *RootArticleRouter) Get() {

	id := this.GetString("id")
	if id != "" {
		this.Data["json"], _ = models.GetArticle(&bson.M{"_id": bson.ObjectIdHex(id)})
		this.ServeJson(true)
	} else {
		limit := common.Webconfig.PageCount
		page, err := strconv.Atoi(this.Ctx.Input.Param(":page"))
		if err != nil {
			page = 1
		}

		article, _, _ := models.GetArticles(&bson.M{}, (page-1)*limit, limit, "-createdtime")

		cat := models.Categories
		this.Data["Category"] = cat
		this.Data["Tags"] = models.Tags
		//nodes := make([]NodeDetail, 0)

		// for _, v := range cat {
		// 	for _, va := range v.Nodes {
		// 		nodes = append(nodes, NodeDetail{
		// 			Name:          va.Name,
		// 			Title:         va.Title,
		// 			Content:       va.Content,
		// 			CategoryTitle: v.Title,
		// 			CreatedTime:   va.CreatedTime,
		// 			UpdatedTime:   va.UpdatedTime,
		// 			Views:         va.Views,
		// 			ArticleTime:   va.ArticleTime,
		// 		})
		// 	}
		// }
		// opt := this.Ctx.Input.Param(":opt")
		// if opt != "" {
		// 	this.Data["notify"] = true
		// }
		// if opt == "update" {
		// 	this.Data["msg"] = "更新成功！"
		// } else if opt == "delete" {
		// 	this.Data["msg"] = "删除成功！"
		// } else if opt == "add" {
		// 	this.Data["msg"] = "添加成功！"
		// }
		this.Data["Articles"] = article
		this.Data["currentitem"] = "article"
		this.Layout = "root/layout.html"
		this.TplNames = "root/article.html"
	}
}

func (this *RootArticleRouter) Post() {
	id := this.GetString("id")
	if len(this.Input()) == 1 { //删除操作
		models.DeleteArticle(&bson.M{"_id": bson.ObjectIdHex(id)})
		this.Data["json"] = true
		this.ServeJson(true)
	} else {
		nodename := this.GetString("selectnode")
		name := this.GetString("name")
		title := this.GetString("title")
		content := this.GetString("content")
		isThumbnail, _ := this.GetBool("isThumbnail")
		featuredPicURL := this.GetString("featuredPicURL")
		tags := this.GetStrings("tags")
		author, _ := this.GetSession("username").(string)
		cat := models.GetCategoryNodeName(nodename)
		if name == "" {
			name = strconv.Itoa(int(bson.Now().UnixNano()))
		}
		if id != "" {
			//更新
			article, _ := models.GetArticle(&bson.M{"_id": bson.ObjectIdHex(id)})
			article.CName = cat.Name
			article.NName = nodename
			article.Name = name
			article.Author = author
			article.Title = title
			article.Tags = tags
			article.FeaturedPicURL = featuredPicURL
			article.ModifiedTime = bson.Now()
			article.Text = template.HTML(content)
			article.IsThumbnail = isThumbnail

			article.SetSummary()
			article.UpdateArticle()
			this.Redirect("/root/article", 302)
		} else {
			//创建
			article := models.Article{
				CName:          cat.Name,
				NName:          nodename,
				Name:           name,
				Author:         author,
				Title:          title,
				Tags:           tags,
				FeaturedPicURL: featuredPicURL,
				CreatedTime:    bson.Now(),
				ModifiedTime:   bson.Now(),
				Text:           template.HTML(content),
				IsThumbnail:    isThumbnail,
			}
			article.SetSummary()
			article.CreatArticle()
			go Publish(&article)
			this.Redirect("/root/article", 302)
		}
	}

}
