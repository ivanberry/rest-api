package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

func init()  {
	cwd, _ := os.Getwd()

	if e := godotenv.Load(cwd + "/config.env"); e != nil {

		fmt.Println(e)
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("user")
	pwd := os.Getenv("password")
	charset := os.Getenv("charset")
	dbName := os.Getenv("db")
	dbHost := os.Getenv("host")

	dbUri := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=%s&parseTime=true", username, pwd, dbHost, dbName, charset)
	fmt.Println(dbUri)

	conn,err := gorm.Open("mysql", dbUri)
	if err != nil {
		log.Fatal("Connect database error")
	}

	db = conn
	db.SingularTable(true)
}

func GetDB() *gorm.DB {
	return db
}