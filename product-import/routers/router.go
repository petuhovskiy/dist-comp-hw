package routers

import (
	"fmt"
	"lib/httputil"
	"lib/pb"
	"log"
	"net/http"
	"product-import/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "product-import/docs" // docs is generated by Swag CLI, you have to import it.
)

func CreateRouter(imp *handlers.Import, authMiddleware func (http.Handler) http.Handler) chi.Router {
	r := chi.NewMux()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		r.Use(authMiddleware)
		r.With(httputil.RequireRole(pb.AuthRole_ADMIN)).Post("/import", imp.Import)
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
