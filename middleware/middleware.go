package middleware

import (
	"log"
	"net/http"
)

type Middleware func(next http.HandlerFunc) http.HandlerFunc

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

func ChainMiddleware(mw ...Middleware) Middleware {
	return func(final http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			last := final
			for i := len(mw) - 1; i >=0; i-- {
				last = mw[i](last)
			}
			last(writer,request)
		}

	}

}

