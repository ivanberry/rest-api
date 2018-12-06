package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ivanberry/rest-api/utils"
	"github.com/jinzhu/gorm"
	"os"
)

type Account struct {
	gorm.Model
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token";sql:"-"`
}

type Token struct {
	UserId uint
	UserName string
	jwt.StandardClaims
}

func (account *Account) Create() (map[string]interface{}) {
	GetDB().Create(account)

	if account.ID <= 0 {
		return utils.Message(false, "Failed to create account, connection error.")
	}

	// Create JWT token for the newly registered account
	tk := &Token{UserId:account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString


	response := utils.Message(true, "用户创建成功")
	response["account"] = account
	return response

}

