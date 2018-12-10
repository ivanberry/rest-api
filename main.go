package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/ivanberry/rest-api/app"
	"github.com/ivanberry/rest-api/controller"
	"github.com/ivanberry/rest-api/middleware"
	"github.com/ivanberry/rest-api/models"
	"log"
	"net/http"
)

var lt = middleware.ChainMiddleware(middleware.WithLogging, middleware.WithTracing)

func main() {

	// db connect may not be this place
	db := models.GetDB()
	defer db.Close()
	db.Debug().DropTable(&models.Post{})
	db.AutoMigrate(&models.Account{}, &models.Post{})

	router := mux.NewRouter()
	router.Use(app.JwtAuthentication)
	router.HandleFunc("/", lt(GetIndex)).Methods("GET")
	router.HandleFunc("/api/user/new", lt(controllers.CreateAccout)).Methods("POST")
	router.HandleFunc("/api/user/login", lt(controllers.Authenticate)).Methods("POST")
	router.HandleFunc("/api/userInfo", lt(controllers.GetUserInfo)).Methods("GET")
	router.HandleFunc("/api/post/new", lt(controllers.CreatePost)).Methods("POST")
	router.HandleFunc("/api/post/{id:[0-9]+}", lt(controllers.GetPost)).Methods("GET")


	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	// Pass
	vars := mux.Vars(r)
	fmt.Printf("vars %v\n", vars)
	fmt.Fprintf(w, "Category: %v\n", vars)
}
