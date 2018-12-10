package models

import (
	"github.com/ivanberry/rest-api/utils"
	"github.com/jinzhu/gorm"
	"strings"
)

// 定义post结构提，用于数据存储
type Post struct {
	gorm.Model
	Title string `json:"title"`
	Content string `json:"content"`
}

func (post *Post) Validate() (map[string]interface{}, bool) {

	if len(post.Title) == 0 {
		return utils.Message(false, "文章标题不能为空！"), false
	}

	if len(post.Content) == 0 {
		return utils.Message(false, "文章内容不能为空！"), false
	}

	if strings.Contains(post.Title, "色情") || strings.Contains(post.Content, "色情") {
		return utils.Message(false, "包含敏感词"), false
	}

	return utils.Message(true, "内容合法！"), true
}

func (post *Post) Create() (map[string]interface{}) {
	if resp, ok := post.Validate(); !ok {
		return resp
	}

	GetDB().Create(post)
	if post.ID <= 0 {
		return utils.Message(false, "创建失败")
	}

	return utils.Message(true, "success")

}

func Get(id uint64) *Post {
	po := &Post{}
	GetDB().Table("post").Where("id = ?", id).First(po)
	if po.Title == "" {
		return nil
	}
	return po
}

