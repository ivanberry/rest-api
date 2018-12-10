package controllers

import (
	"encoding/json"
	"github.com/ivanberry/rest-api/models"
	"github.com/ivanberry/rest-api/utils"
	"net/http"
	"strconv"
	"strings"
)

var CreatePost = func(w http.ResponseWriter, r *http.Request) {
	post := &models.Post{}

	// 解析r.Body
	if err := json.NewDecoder(r.Body).Decode(post); err != nil {
		utils.Respond(w, utils.Message(false, "参数错误!"))
		return
	}

	// Store in DB, and populate into post
	resp := post.Create()
	utils.Respond(w, resp)
}

var GetPost = func(w http.ResponseWriter, r *http.Request) {

	// 通过路由参数获取 /api/post/{id}
	postIdStr := strings.Split(r.URL.Path, "/")[3]
	postId, err := strconv.ParseUint(postIdStr, 10, 32 )
	if err != nil {
		utils.Message(false, "参数错误!")
	}

	post := models.Get(postId)

	if post == nil {
		utils.Message(false, "文章不存在")
	}

	resp := utils.Message(true, "success")
	resp["post"] = post
	utils.Respond(w, resp)

}