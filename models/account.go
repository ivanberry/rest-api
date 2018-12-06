package models

import "github.com/dgrijalva/jwt-go"

type Account struct {
	Name string
}

type Token struct {
	UserId uint
	UserName string
	jwt.StandardClaims
}