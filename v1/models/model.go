package models

import (
	"html/template"
	"strings"
	"time"
)

import (
	. "../common"
)

import (
	"labix.org/v2/mgo/bson"
)

type ArticleMetadata struct {
	Name           string
	Author         string
	IsPage         bool
	Title          string
	Tags           []string
	FeaturedPicURL string
	Summary        string
	CreatedTime    time.Time
	ModifiedTime   time.Time
	Hits           int64
}

type Article struct {
	Id_      bson.ObjectId `bson:"_id"`
	Metadata ArticleMetadata
	Text     template.HTML
	Comments []*Comment
}

func (meta *ArticleMetadata) CreatedTimeRFC3339() string {
	return meta.CreatedTime.String()
}

func (meta *ArticleMetadata) ModifiedTimeRFC3339() string {
	return meta.ModifiedTime.String()
}

func (meta *ArticleMetadata) CreatedTimeHumanReading() string {
	return meta.CreatedTime.String()
}

func (meta *ArticleMetadata) GetCreatedTime() time.Time {
	return meta.CreatedTime.Local()
}

func (meta *ArticleMetadata) GetShortMonth(t1 time.Time) string {
	return t1.Month().String()[0:3]
}

func (meta *ArticleMetadata) ModifiedTimeHumanReading() string {
	return meta.ModifiedTime.String()
}

func (meta *ArticleMetadata) TagRawList() string {
	return strings.Join(meta.Tags, ", ")
}

func (meta *ArticleMetadata) HasFeaturedPic() bool {
	if len(meta.FeaturedPicURL) == 0 {
		return false
	}
	return true
}

func (meta *ArticleMetadata) HasSummary() bool {
	if len(meta.Summary) == 0 {
		return false
	}
	return true
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

func (m *CommentMetadata) BuildFromJson(json interface{}) {
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

func (meta *CommentMetadata) CreatedTimeHumanReading() string {
	return TimeHumanReading(meta.CreatedTime)
}

type Comment struct {
	Metadata CommentMetadata
	Text     template.HTML
}

type Category struct {
	Id_         bson.ObjectId `bson:"_id"`
	Title       string
	Content     string
	CreatedTime bson.Time
	UpdatedTime bson.Time
	Views       int64
	NodeTime    bson.Time
	NodeCount   int64
	NodeIds     []bson.ObjectId
}

type Node struct {
	Id_          bson.ObjectId `bson:"_id"`
	CId_         bson.ObjectId `bson:"_cid"`
	Title        string
	Content      string
	CreatedTime  bson.Time
	UpdatedTime  bson.Time
	Views        int64
	TopicTime    bson.Time
	ArticleCount int64
}

type TagWrapper struct {
	Id_          bson.ObjectId `bson:"_id"`
	Name         string
	Count        int
	CreatedTime  bson.Time
	ModifiedTime bson.Time
	ArticleIds   []bson.ObjectId
}
