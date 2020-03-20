package routers

import (
	"auth/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"


	httpSwagger "github.com/swaggo/http-swagger"

	_ "auth/docs" // docs is generated by Swag CLI, you have to import it.
)

func CreateRouter(authV1 *handlers.Auth) chi.Router {
	r := chi.NewMux()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/v1", func(r chi.Router) {
		r.Post("/signup", authV1.Signup)
		r.Post("/signin", authV1.Signin)
		r.Post("/refresh", authV1.Refresh)
		r.Post("/validate", authV1.Validate)
	})

	log.Println("Attached swagger ui at /swagger")

	r.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/", http.StatusMovedPermanently)
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The url pointing to API definition"
	))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `<a href="/swagger">Swagger UI</a>`)
	})

	return r
}