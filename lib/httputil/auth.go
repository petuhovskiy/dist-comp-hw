package httputil

import (
	"context"
	"github.com/go-chi/render"
	"lib/pb"
	"log"
	"net/http"
	"strings"
)

type ctxKey string

const (
	roleKey ctxKey = "role"
)

func AuthMiddleware(cli pb.AuthClient) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			authHeader = strings.TrimPrefix(authHeader, "Bearer ")

			info, err := cli.Validate(context.Background(), &pb.ValidateRequest{
				AccessToken: authHeader,
			})
			if err != nil {
				log.Println("failed to validate token, err=", err)
				render.Render(w, r, ErrUnauthorized)
				return
			}
			log.Println("Auth success, info=", info)

			ctx := r.Context()
			ctx = context.WithValue(ctx, roleKey, info.GetRole())

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RequireRole(required pb.AuthRole) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			role := r.Context().Value(roleKey)

			if role != required {
				log.Println("bad role, role=", role)
				render.Render(w, r, ErrUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
