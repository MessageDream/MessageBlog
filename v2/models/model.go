package models

import (
	"html/template"
	//"strings"
	//"log"
	"regexp"
	"strings"
	"time"
)

import (
	. "../common"
)

import (
	"labix.org/v2/mgo/bson"
)

type Article struct {
	Id_            bson.ObjectId `bson:"_id"`
	CName          string
	NName          string
	Name           string
	Author         string
	Title          string
	Text           template.HTML
	Tags           []string
	FeaturedPicURL string
	Summary        template.HTML
	Views          int
	Comments       []*Comment
	IsThumbnail    bool
	CreatedTime    time.Time
	ModifiedTime   time.Time
}

func (article *Article) SetSummary() {
	if article.IsThumbnail {
		article.Summary = article.Text
	} else {
		strs := strings.Split(string(article.Text), "<!--more-->")
		//beego.Error(strs[0])
		n := len(strs)
		if n > 0 {
			article.Summary = template.HTML(strs[0])
		}
	}
}

func (article *Article) GetFirstParagraph() *template.HTML {
	rx := regexp.MustCompile(`<p>(.*)</p>`)
	p := rx.FindStringSubmatch(string(article.Text))
	//beego.Error(p)
	n := len(p)
	if n > 1 {
		rep := template.HTML(p[1] + "...")
		return &rep
	}
	return nil
}

func (article *Article) GetCategory() *Category {
	// c := DB.C("category")
	// var category Category
	// c.Find(bson.M{"name": article.CName}).One(&category)
	// return &category
	var category Category
	for _, v := range Categories {
		if v.Name == article.CName {
			category = v
			break
		}
	}
	return &category
}

func (article *Article) GetNode() *Node {
	var node Node
	for _, v := range Categories {
		if v.Name == article.CName {
			for _, va := range v.Nodes {
				if va.Name == article.NName {
					node = va
					break
				}
			}
			break
		}
	}

	return &node
}

func (article *Article) GetTags() *[]TagWrapper {
	return article.GetSelfTags()
}

func (article *Article) CreatArticle() error {
	article.Id_ = bson.NewObjectId()
	c := DB.C("article")
	err := c.Insert(article)
	go setTags(&article.Tags, article.Id_)
	return err
}

func (article *Article) UpdateArticle() error {
	c := DB.C("article")
	err := c.UpdateId(article.Id_, article)
	go setTags(&article.Tags, article.Id_)

	return err
}

func (article *Article) GetCommentCount() int {
	return 1
}

func (article *Article) GetAroundArticle() (*Article, *Article, error) {
	c := DB.C("article")
	var preresult, nextresult Article
	err := c.Find(&bson.M{"createdtime": &bson.M{"$lt": article.CreatedTime}}).Sort("-createdtime").Limit(1).One(&preresult)

	err = c.Find(&bson.M{"createdtime": &bson.M{"$gt": article.CreatedTime}}).Sort("createdtime").Limit(1).One(&nextresult)
	return &preresult, &nextresult, err
}

func (article *Article) GetSameTagArticles(limit int) (articles []Article) {
	ids := make([]bson.ObjectId, 0)
	for _, v := range Tags {
		for _, tag := range article.Tags {
			if tag == v.Title || tag == v.Name {
				for _, va := range v.ArticleIds {
					if va != article.Id_ {
						ids = append(ids, va)
					}
				}
			}
		}
	}
	d := DB.C("article")
	d.Find(&bson.M{"_id": &bson.M{"$in": ids}}).Limit(limit).All(&articles)
	return
}

func (article *Article) GetSelfTags() *[]TagWrapper {
	var tags []TagWrapper
	for _, v := range Tags {
		for _, va := range v.ArticleIds {
			if va != article.Id_ {
				tags = append(tags, v)
			}
		}
	}
	return &tags
}

func (article *Article) HasFeaturedPic() bool {
	if len(article.FeaturedPicURL) == 0 {
		return false
	}
	return true
}

func (article *Article) HasSummary() bool {
	if len(article.Summary) == 0 {
		return false
	}
	return true
}

func (article *Article) UpdateViews() {
	article.Views ++
	article.UpdateArticle()
}

type CommentIndexItem struct {
	Name         string
	CommentNames []string
}

type CommentMetadata struct {
	Name        string
	Author      string
	ArticleName string
	UAgent      string
	URL         string
	IP          string
	Email       string
	EmailHash   string
	CreatedTime int64
}

func (m *Comment) BuildFromJson(json interface{}) {
	var jsonMap map[string]interface{}
	jsonMap = json.(map[string]interface{})
	for k, v := range jsonMap {
		switch vv := v.(type) {
		case string:
			str := vv
			switch k {
			case "Name":
				m.Name = str
			case "Author":
				m.Author = str
			case "URL":
				m.URL = str
			case "IP":
				m.IP = str
			case "Email":
				m.Email = str
			case "EmailHash":
				m.EmailHash = str
			case "UAgent":
				m.UAgent = str
			case "ArticleName":
				m.ArticleName = str
			}
		case float64:
			if k == "CreatedTime" {
				m.CreatedTime = int64(vv)
			}
		default:
		}
	}
}

func (meta *Comment) CreatedTimeHumanReading() string {
	return TimeHumanReading(meta.CreatedTime)
}

type Comment struct {
	CommentMetadata
	Text template.HTML
}

type Node struct {
	Name        string
	Title       string
	Content     string
	CreatedTime time.Time
	UpdatedTime time.Time
	Views       int64
	ArticleTime time.Time
}

func (node *Node) GetAllArticles(offset int, limit int, sort string) (*[]Article, int, error) {
	c := DB.C("article")

	var article []Article
	q := bson.M{"nname": node.Name}
	query := c.Find(q).Skip(offset).Limit(limit)
	if sort != "" {
		query = query.Sort(sort)
	}
	err := query.All(&article)
	total, _ := c.Find(q).Count()
	return &article, total, err
}

func (node *Node) GetArticleCount() (int, error) {
	c := DB.C("article")
	return c.Find(bson.M{"nname": node.Name}).Count()
}

// func (node *Node) GetCategory() (*Category, error) {
// 	c := DB.C("category")
// 	var category Category
// 	err := c.Find(bson.M{"_id": node.CId_}).One(&category)
// 	return &category, err
// }

// func (node *Node) CreatNode() error {
// 	//node.Id_ = bson.NewObjectId()
// 	c := DB.C("node")
// 	err := c.Insert(node)
// 	return err
// }

// func (node *Node) UpdateNode() error {
// 	c := DB.C("node")
// 	err := c.UpdateId(node.Id_, node)
// 	return err
// }

type Category struct {
	Id_         bson.ObjectId `bson:"_id"`
	Name        string
	Title       string
	Content     string
	CreatedTime time.Time
	UpdatedTime time.Time
	Views       int
	NodeTime    time.Time
	Nodes       []Node
}

func (category *Category) CreatCategory() error {
	//category.Id_ = bson.NewObjectId()
	c := DB.C("category")
	err := c.Insert(category)
	SetAppCategories()
	return err
}

func (category *Category) UpdateCategory() error {
	c := DB.C("category")
	err := c.UpdateId(category.Id_, category)
	SetAppCategories()
	return err
}

// func (category *Category) GetAllNodes() *[]Node {
// 	c := DB.C("node")
// 	var nodes []Node
// 	c.Find(&bson.M{"_cid": category.Id_}).All(&nodes)
// 	return &nodes
// }

// func (category *Category) SetNodeId(nid bson.ObjectId) {
// 	if category.NodeIds != nil {
// 		category.NodeIds = append(category.NodeIds, nid)
// 		removeDuplicate(&category.NodeIds)
// 	} else {
// 		category.NodeIds = []bson.ObjectId{nid}
// 	}
// 	category.NodeCount = len(category.NodeIds)
// }

type TagWrapper struct {
	Id_          bson.ObjectId `bson:"_id"`
	Name         string
	Title        string
	Count        int
	CreatedTime  time.Time
	ModifiedTime time.Time
	ArticleIds   []bson.ObjectId
}

func (tag *TagWrapper) SetTag() error {
	c := DB.C("tags")
	var err error
	flag := false
	for _, v := range Tags {
		if tag.Name == v.Name {
			v.ArticleIds = append(v.ArticleIds, tag.ArticleIds...)
			removeDuplicate(&v.ArticleIds)
			v.Count = len(v.ArticleIds)
			v.ModifiedTime = bson.Now()
			err = c.UpdateId(v.Id_, v)
			flag = true
			break
		}
	}

	if !flag {
		tag.Id_ = bson.NewObjectId()
		tag.CreatedTime = bson.Now()
		tag.ModifiedTime = bson.Now()
		Tags = append(Tags, *tag)
		err = c.Insert(tag)
	}

	SetAppTags()
	return err
}

type Subscription struct {
	Id_    bson.ObjectId `bson:"_id"`
	Email  string
	Uid    string
	Status bool
}

func (subscription *Subscription) Set() error {
	c := DB.C("subscription")
	var sub Subscription
	err := c.Find(bson.M{"email": subscription.Email}).One(&sub)
	if err != nil {
		subscription.Id_ = bson.NewObjectId()
		err = c.Insert(subscription)
	} else {
		sub.Email = subscription.Email
		sub.Status = subscription.Status
		sub.Uid = subscription.Uid
		err = c.UpdateId(sub.Id_, sub)
	}
	return err
}

func (subscription *Subscription) UpdateState() error {
	c := DB.C("subscription")
	var sub Subscription
	err := c.Find(bson.M{"uid": subscription.Uid}).One(&sub)
	if err != nil {
		return err
	}
	sub.Status = subscription.Status
	err = c.UpdateId(sub.Id_, sub)
	return err
}
