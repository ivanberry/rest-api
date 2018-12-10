package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ivanberry/rest-api/utils"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
	"time"
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

	// validate email
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

	//Don't store the raw password, but the hashed one
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


/**
Handle login
0. get request body
1. compare password
2. return the particular resp
 */
func Login(email, password string) (map[string]interface{})  {

	account := &Account{}

	//fetch data from db with email
	//and store the value to account if no err
	err := GetDB().Table("account").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Message(false, "用户不存在")
		}
		return utils.Message(false, "链接错误")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil {
		return utils.Message(false, "账户或密码错误")
	}

	account.Password = ""

	exp := time.Now().Add(time.Hour * 72).Unix()

	// Login in success and return account info with jwt token
	tk := &Token{account.ID, jwt.StandardClaims{
		ExpiresAt: exp,
	}}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))

	resp := utils.Message(true,"success")
	resp["account"] = account
	resp["token"] = tokenString
	return resp
}


/**
get User infos with id
 */
func Getuser(u uint) *Account  {
	account := &Account{}
	GetDB().Table("account").Where("id = ?", u).First(account)
	if account.Email == "" {
		return nil
	}
	account.Password = ""
	return account
}
