package middleware

import (
	"log"
	"net/http"
)

type middleware func(next http.HandlerFunc) http.HandlerFunc

func WithLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("Logged connect from %s", request.RemoteAddr)
		next.ServeHTTP(writer, request)
	}
}

func WithTracing(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("Tracing request for %s", request.RequestURI)
		next.ServeHTTP(writer, request)
	}
}

