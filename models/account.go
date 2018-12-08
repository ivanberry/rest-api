package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ivanberry/rest-api/utils"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
)

type Account struct {
	gorm.Model
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token";sql:"-"`
}

type Token struct {
	UserId uint
	jwt.StandardClaims
}

func (account *Account) Validate() (map[string]interface{}, bool) {

	// validate emial
	if !strings.Contains(account.Email, "@") {
		return utils.Message(false, "请输入邮箱地址."), false
	}

	if len(account.Password) < 6 {
		return utils.Message(false, "密码太弱."), false
	}

	temp := &Account{}

	// check email duplicate
	// 通过查询将获取的数据放置在了temp中?
	err := GetDB().Table("account").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return utils.Message(false, "链接错误，请重试!"), false
	}

	if temp.Email != "" {
		return utils.Message(false, "此邮箱已通过注册."), false
	}

	return utils.Message(true, "验证通过"), true

}
	

func (account *Account) Create() (map[string]interface{}) {

	if resp, ok := account.Validate(); !ok {
		return resp
	}

	//Don't store the raw password, but the hash on
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashPassword)

	//TODO: how to store the jwt token
	GetDB().Create(account)

	if account.ID <= 0 {
		return utils.Message(false, "Failed to create account, connection error.")
	}

	// Create JWT token for the newly registered account
	tk := &Token{UserId:account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	//TODO: auto ignore nil value
	account.Password = ""

	response := utils.Message(true, "用户创建成功")
	response["account"] = account
	return response
}

