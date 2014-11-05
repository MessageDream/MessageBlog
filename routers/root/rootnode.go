package root

import (
	//"log"
	"strconv"
	"time"
)

import (
	"../../models"
)

import (
	"labix.org/v2/mgo/bson"
)

type NodeDetail struct {
	Name          string
	Title         string
	Content       string
	CategoryTitle string
	CategoryId    string
	CreatedTime   time.Time
	UpdatedTime   time.Time
	Views         int64
	ArticleTime   time.Time
}

type RootNodeRouter struct {
	rootBaseRouter
}

func (this *RootNodeRouter) Get() {
	name := this.GetString("name")
	cat := models.Categories
	if name != "" {
		var nodeDetail NodeDetail
		for _, v := range cat {
			flag := false
			for _, va := range v.Nodes {
				if va.Name == name {
					nodeDetail = NodeDetail{
						Name:          va.Name,
						Title:         va.Title,
						Content:       va.Content,
						CategoryTitle: v.Title,
						CategoryId:    v.Id_.Hex(),
						CreatedTime:   va.CreatedTime,
						UpdatedTime:   va.UpdatedTime,
						Views:         va.Views,
						ArticleTime:   va.ArticleTime,
					}
					flag = true
					break
				}
			}
			if flag {
				break
			}
		}
		this.Data["json"] = nodeDetail
		this.ServeJson(true)
	} else {

		this.Data["Category"] = cat
		nodes := make([]NodeDetail, 0)

		for _, v := range cat {
			for _, va := range v.Nodes {
				nodes = append(nodes, NodeDetail{
					Name:          va.Name,
					Title:         va.Title,
					Content:       va.Content,
					CategoryTitle: v.Title,
					CategoryId:    v.Id_.Hex(),
					CreatedTime:   va.CreatedTime,
					UpdatedTime:   va.UpdatedTime,
					Views:         va.Views,
					ArticleTime:   va.ArticleTime,
				})
			}
		}
		this.Data["Nodes"] = nodes
		this.Data["currentitem"] = "node"
		this.Layout = "root/layout.html"
		this.TplNames = "root/node.html"
	}
}

func (this *RootNodeRouter) Post() {
	if len(this.Input()) == 2 { //删除操作
		cid := this.GetString("id")
		nname := this.GetString("nname")
		for _, v := range models.Categories {
			if v.Id_.Hex() == cid {
				for in, va := range v.Nodes {
					if va.Name == nname {
						v.Nodes = append(v.Nodes[:in], v.Nodes[(in+1):]...)
						break
					}
				}
				v.UpdateCategory()
				break
			}
		}
		this.Data["json"] = true
		this.ServeJson(true)
	} else {
		categoryid := this.GetString("selectcategory")
		name := this.GetString("name")
		title := this.GetString("title")
		content := this.GetString("content")
		if name == "" {
			name = strconv.Itoa(int(bson.Now().UnixNano()))
		}

		for _, v := range models.Categories {
			if v.Id_.Hex() == categoryid {
				flag := false
				for _, va := range v.Nodes {
					if va.Name == name { //更新
						va.Title = title
						va.Content = content
						va.UpdatedTime = bson.Now()
						flag = true
						break
					}
				}

				if !flag { //添加
					node := models.Node{
						Name:        name,
						Title:       title,
						Content:     content,
						CreatedTime: bson.Now(),
						UpdatedTime: bson.Now(),
					}
					v.Nodes = append(v.Nodes, node)
				}

				v.UpdateCategory()
				break
			}
		}
		this.Redirect("/root/node", 302)
	}
}
