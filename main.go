package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ivanberry/rest-api/middleware"
	"github.com/ivanberry/rest-api/models"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

var lt = middleware.ChainMiddleware(middleware.WithLogging, middleware.WithTracing)

func main() {
	dbUri := models.GetDB()
	db, err := gorm.Open("mysql", dbUri)
	if err != nil {
		log.Fatal("Connect database error")
	}

	defer db.Close()
	db.AutoMigrate(&models.Account{}, &models.Contact{})

	router := mux.NewRouter()
	router.HandleFunc("/", lt(GetIndex)).Methods("GET")
	router.HandleFunc("/index", middleware.WithLogging(middleware.WithTracing(GetIndex))).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	// Pass
	vars := mux.Vars(r)
	fmt.Printf("vars %v\n", vars)
	fmt.Fprintf(w, "Category: %v\n", vars)
}
