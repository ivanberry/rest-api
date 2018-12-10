package models

import (
	"github.com/ivanberry/rest-api/utils"
	"github.com/jinzhu/gorm"
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

