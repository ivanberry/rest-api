package app

import (
	"github.com/ivanberry/rest-api/utils"
	"net/http"
	"strings"
)

var JwtAuthentication = func(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		noNeedAuth := []string{"/api/user/new", "/api/user/login"}
		requestPath := r.URL.Path

		for _, value := range noNeedAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
			}
		}

		response := make(map[string] interface{})
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			response = utils.Message(false, "Invalid auth token header.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "appllication/json")
			utils.Respond(w, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) !=2 {
			response = utils.Message(false, "Invalid auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
		}

		tokenPart := splitted[1]
		//tk := &models.Token{}
	})
}

