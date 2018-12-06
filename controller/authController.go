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
