package models

import (
	"fmt"
	"os"
)

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var (
	DB *mgo.Database
)

func InitDb(conn string) {
	if conn == "" {
		fmt.Println("数据库地址还没有配置,请到config内配置db字段.")
		os.Exit(1)
	}

	session, err := mgo.Dial(conn)
	if err != nil {
		fmt.Println("MongoDB连接失败:", err.Error())
		os.Exit(1)
	}

	session.SetMode(mgo.Monotonic, true)

	DB = session.DB("messageblog")

}

func GetArticles(condition *bson.M, offset int, limit int, sort string) (*[]Article, int, error) {
	c := DB.C("article")

	var article []Article
	query := c.Find(condition).Skip(offset).Limit(limit)
	if sort != "" {
		query = query.Sort(sort)
	}
	err := query.All(&article)
	total, _ := c.Find(condition).Count()
	return &article, total, err
}

func GetArticlesByTag(condition *bson.M, offset int, limit int, sort string) (*[]Article, int, error) {
	c := DB.C("tags")
	var tag TagWrapper
	err := c.Find(condition).One(&tag)
	result, total, _ := GetArticles(&bson.M{"_id": bson.M{"$in": tag.ArticleIds}}, offset, limit, sort)
	return result, total, err
}

func GetArticleCount() int {
	c := DB.C("article")
	total, _ := c.Count()
	return total
}

func GetArticle(condition *bson.M) (*Article, error) {
	c := DB.C("article")
	var article Article
	err := c.Find(condition).One(&article)
	return &article, err
}

func GetAroundArticleName(precondition, nextcondition *bson.M) (string, string, error) {
	c := DB.C("article")
	var preresult, nextresult map[string]map[string]string
	err := c.Find(precondition).Sort("-metadata.createdtime").Limit(1).Select(bson.M{"metadata.name": 1, "_id": 0}).One(&preresult)

	err = c.Find(nextcondition).Sort("metadata.createdtime").Limit(1).Select(bson.M{"metadata.name": 1, "_id": 0}).One(&nextresult)
	return preresult["metadata"]["name"], nextresult["metadata"]["name"], err
}

func CreatArticle(article *Article) error {
	article.Id_ = bson.NewObjectId()
	c := DB.C("article")
	err := c.Insert(article)
	go setTags(&article.Metadata.Tags, article.Id_)
	return err
}

func UpdateArticle(article *Article) error {
	c := DB.C("article")
	err := c.UpdateId(article.Id_, article)
	go setTags(&article.Metadata.Tags, article.Id_)

	return err
}

func DeleteArticle(condition *bson.M) (*mgo.ChangeInfo, error) {
	c := DB.C("article")
	return c.RemoveAll(condition)
}

func GetTags(condition *bson.M, offset int, limit int, sort string) (*[]TagWrapper, int, error) {
	c := DB.C("tags")
	var tags []TagWrapper
	query := c.Find(condition).Skip(offset).Limit(limit)
	if sort != "" {
		query = query.Sort(sort)
	}
	err := query.All(&tags)
	total, _ := c.Find(condition).Count()

	return &tags, total, err
}

func GetTagCount() int {
	c := DB.C("tags")
	total, _ := c.Count()
	return total
}

func SetTag(tag *TagWrapper) error {
	c := DB.C("tags")
	var oldtag TagWrapper
	err := c.Find(bson.M{"name": tag.Name}).One(&oldtag)
	if err != nil && err.Error() == "not found" {
		tag.Id_ = bson.NewObjectId()
		tag.CreatedTime = bson.Now()
		tag.ModifiedTime = bson.Now()
		c.Insert(tag)
	} else if err != nil {
		return err
	} else {
		oldtag.ArticleIds = append(oldtag.ArticleIds, tag.ArticleIds...)
		removeDuplicate(&oldtag.ArticleIds)
		oldtag.Count = len(oldtag.ArticleIds)
		oldtag.ModifiedTime = bson.Now()
		c.UpdateId(oldtag.Id_, oldtag)
	}
	return nil
}

func removeDuplicate(slis *[]bson.ObjectId) {
	found := make(map[bson.ObjectId]bool)
	j := 0
	for i, val := range *slis {
		if _, ok := found[val]; !ok {
			found[val] = true
			(*slis)[j] = (*slis)[i]
			j++
		}
	}
	*slis = (*slis)[:j]
}

func setTags(tags *[]string, aid bson.ObjectId) {
	for _, v := range *tags {
		tag := &TagWrapper{
			Name:       v,
			Count:      1,
			ArticleIds: []bson.ObjectId{aid},
		}
		SetTag(tag)
	}
}
