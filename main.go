//go:generate swag init
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v4"

	"github.com/petuhovskiy/dist-comp-hw/config"
	"github.com/petuhovskiy/dist-comp-hw/handlers"
	"github.com/petuhovskiy/dist-comp-hw/repos/psql"
	"github.com/petuhovskiy/dist-comp-hw/routers"
	"github.com/petuhovskiy/dist-comp-hw/service"
)

// @title Internet Shop API
// @version 1.0
// @description This is sample internet shop, made as dist-comp homework.

// @contact.name Arthur Petukhovsky
// @contact.url https://t.me/petuhovskiy
// @contact.email petuhovskiy@yandex.ru

// @host localhost:8080
// @BasePath /
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf, err := config.EnvConfig()
	if err != nil {
		log.Fatal("failed to read config, err=", err)
	}

	// db
	conn, err := pgx.Connect(context.Background(), conf.PostgresAddr)
	if err != nil {
		log.Fatal("failed to connect to postgres, err=", err)
	}
	defer conn.Close(context.Background())

	// repositories
	productsRepo := psql.NewProducts(conn)

	// applying migrations
	err = productsRepo.Migrate()
	if err != nil {
		log.Fatal("failed to migrate products, err=", err)
	}

	// initializing services
	productsService := service.NewProducts(productsRepo)

	// initializing handlers
	productsV1 := handlers.NewProductsV1(productsService)
	handler := routers.CreateRouter(productsV1)

	log.Println("Serving at http://localhost" + conf.BindAddr)
	err = http.ListenAndServe(conf.BindAddr, handler)
	log.Fatal(err)
}
