package handlers

import (
	"app/auth"
	"github.com/go-chi/render"
	"log"
	"net/http"
	"strings"
)

func AuthMiddleware(cli *auth.Client) func (handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			authHeader = strings.TrimPrefix(authHeader, "Bearer ")

			info, err := cli.Validate(authHeader)
			if err != nil {
				log.Println("failed to validate token, err=", err)
				render.Render(w, r, ErrUnauthorized)
				return
			}
			log.Println("Auth success, info=", info)

			next.ServeHTTP(w, r)
		})
	}
}