//go:generate swag init
package main

import (
	"app/batchimport"
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"lib/httputil"
	"lib/pb"
	"log"
	"net/http"

	"github.com/jackc/pgx/v4"

	"app/config"
	"app/handlers"
	"app/repos/psql"
	"app/routers"
	"app/service"
)

// @title Internet Shop API
// @version 1.0
// @description This is sample internet shop, made as dist-comp homework.

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @contact.name Arthur Petukhovsky
// @contact.url https://t.me/petuhovskiy
// @contact.email petuhovskiy@yandex.ru

// @host localhost:8080
// @BasePath /
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf, err := config.EnvConfig()
	if err != nil {
		log.Panic("failed to read config, err=", err)
	}

	// db
	conn, err := pgx.Connect(context.Background(), conf.PostgresAddr)
	if err != nil {
		log.Panic("failed to connect to postgres, err=", err)
	}
	defer conn.Close(context.Background())

	// clients gRPC
	gconn, err := grpc.Dial(
		conf.AuthGrpc,
		grpc.WithInsecure(),
	)
	if err != nil {
		logrus.Fatal(err)
	}
	authCli := pb.NewAuthClient(gconn)

	// repositories
	productsRepo := psql.NewProducts(conn)

	// applying migrations
	err = productsRepo.Migrate()
	if err != nil {
		log.Panic("failed to migrate products, err=", err)
	}

	// initializing services
	productsService := service.NewProducts(productsRepo)

	// initializing handlers
	productsV1 := handlers.NewProductsV1(productsService)
	authMiddleware := httputil.AuthMiddleware(authCli)
	handler := routers.CreateRouter(productsV1, authMiddleware)

	batchImport := batchimport.NewImportWatcher(conf.AmqpURL, conf.QueueImport, productsRepo)
	go batchImport.Start(context.Background())

	log.Println("Serving at http://localhost" + conf.BindAddr)
	err = http.ListenAndServe(conf.BindAddr, handler)
	log.Println(err)
}
