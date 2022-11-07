package middlewares

import (
	"net/http"
	"restapiexample/src/helpers"
	"strings"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		headerToken := r.Header.Get("Authorization")
			if !strings.Contains(headerToken, "Bearer") {
				helpers.New(401, true, "invalid token", nil).Send(w)
				return
			}

			token := strings.Replace(headerToken, "Bearer ", "", -1)
			parseToken, err := helpers.ParseToken(token)
			if err != nil {
				helpers.New(401, true, "invalid token", err.Error()).Send(w)
				return
			}

			if parseToken.Role !="admin" {
				helpers.New(401, true, "you don't have access to this feature", nil).Send(w)
				return
			}

			if parseToken == nil {
				helpers.New(401, true, "invalid token", nil).Send(w)
				return
			}
			next.ServeHTTP(w, r)
		}
	}

func IsAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		headerToken := r.Header.Get("Authorization")

		if !strings.Contains(headerToken, "Bearer") {
			helpers.New(401, true, "invalid token", nil).Send(w)
			return
		}

		token := strings.Replace(headerToken, "Bearer ","", -1)
		parseToken, err := helpers.ParseToken(token)
		if err != nil {
			helpers.New(401, true, "invalid token", err.Error()).Send(w)
			return
		}

		if parseToken.Role !="admin" {
			helpers.New(401, true, "you don't have access to this feature", nil).Send(w)
			return
		}

		if parseToken == nil {
			helpers.New(401, true, "invalid token", nil).Send(w)
			return
		}

		next.ServeHTTP(w, r)
	}
}
