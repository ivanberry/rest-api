package controllers

import (
	"encoding/json"
	"github.com/ivanberry/rest-api/models"
	"github.com/ivanberry/rest-api/utils"
	"net/http"
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