package controllers

import (
	"encoding/json"
	"github.com/ivanberry/rest-api/models"
	"github.com/ivanberry/rest-api/utils"
	"net/http"
)

var CreateAccout = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}

	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		utils.Respond(w, utils.Message(false, "传入参数不对"))
		return
	}

	resp := account.Create()
	utils.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	// model variable to populate
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		utils.Respond(w, utils.Message(false, "入参错误"))
		return
	}

	// login
	resp := models.Login(account.Email, account.Password)
	//w.Header().Add("Authenticate", "Bear " + resp.account)
	utils.Respond(w, resp)
}

var GetUserInfo = func(w http.ResponseWriter, r *http.Request) {

	userId := r.Context().Value("user").(uint);
	acc := models.Getuser(userId)

	if acc == nil {
		utils.Message(false, "用户不存在")
	}

	resp := utils.Message(true, "success")
	resp["user"] = acc;
	utils.Respond(w, resp)
}
