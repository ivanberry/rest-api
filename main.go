package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/ivanberry/rest-api/app"
	"github.com/ivanberry/rest-api/controller"
	"github.com/ivanberry/rest-api/middleware"
	"github.com/ivanberry/rest-api/models"
	"github.com/rs/cors"
	"log"
	"net/http"
)

var lt = middleware.ChainMiddleware(middleware.WithLogging, middleware.WithTracing)

func main() {

	// db connect may not be this place
	db := models.GetDB()
	defer db.Close()
	db.Debug().DropTable(&models.Account{})
	db.AutoMigrate(&models.Account{}, &models.Contact{})

	router := mux.NewRouter()
	router.Use(app.JwtAuthentication)
	router.HandleFunc("/", lt(GetIndex)).Methods("GET")
	router.HandleFunc("/api/user/new", lt(controllers.CreateAccout)).Methods("POST")
	router.HandleFunc("/api/user/login", lt(controllers.Authenticate)).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	routerWithCorsAll := c.Handler(router)


	log.Fatal(http.ListenAndServe(":8000", routerWithCorsAll))
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	// Pass
	vars := mux.Vars(r)
	fmt.Printf("vars %v\n", vars)
	fmt.Fprintf(w, "Category: %v\n", vars)
}
