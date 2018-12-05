package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ivanberry/rest-api/middleware"
	"log"
	"net/http"
)

var lt = middleware.ChainMiddleware(middleware.WithLogging, middleware.WithTracing)

func main() {
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
