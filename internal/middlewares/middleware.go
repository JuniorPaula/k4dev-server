package middlewares

import (
	"knowledge-api/internal/auth"
	"knowledge-api/internal/utils"
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func Authenticated(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.TokenValidate(r); err != nil {
			utils.ErrorJSON(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
