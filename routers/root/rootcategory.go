package root

import (
	"strconv"
)

import (
	"../../models"
)

import (
	"gopkg.in/mgo.v2/bson"
)

type RootCategoryRouter struct {
	rootBaseRouter
}

func (this *RootCategoryRouter) Get() {
	id := this.GetString("id")
	if id != "" {
		for _, v := range models.Categories {
			if v.Id_.Hex() == id {
				this.Data["json"] = v
				break
			}
		}
		this.ServeJson(true)
	} else {
		this.Data["Category"] = models.Categories
		this.Data["currentitem"] = "category"
		this.Layout = "root/layout.html"
		this.TplNames = "root/category.html"
	}
}

func (this *RootCategoryRouter) Post() {
	id := this.GetString("id")
	if len(this.Input()) == 1 { //删除操作
		models.DeleteCategory(&bson.M{"_id": bson.ObjectIdHex(id)})
		this.Data["json"] = true
		this.ServeJson(true)
	} else {
		name := this.GetString("name")
		title := this.GetString("title")
		content := this.GetString("content")
		if name == "" {
			name = strconv.Itoa(int(bson.Now().UnixNano()))
		}
		if id != "" {
			for _, v := range models.Categories {
				if v.Id_.Hex() == id {
					v.Name = name
					v.Title = title
					v.Content = content
					v.UpdatedTime = bson.Now()
					v.UpdateCategory()
					break
				}
			}
		} else {
			cat := models.Category{
				Id_:         bson.NewObjectId(),
				Name:        name,
				Title:       title,
				Content:     content,
				CreatedTime: bson.Now(),
				UpdatedTime: bson.Now(),
				NodeTime:    bson.Now(),
			}
			cat.CreatCategory()
		}

		this.Redirect("/root/category", 302)
	}
}
