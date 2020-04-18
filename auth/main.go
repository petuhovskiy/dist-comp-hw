//go:generate swag init
package main

import (
	"auth/grpcserv"
	"context"
	"github.com/cristalhq/jwt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"lib/pb"
	"log"
	"net/http"

	"github.com/jackc/pgx/v4"

	"auth/config"
	"auth/handlers"
	"auth/repos/psql"
	"auth/routers"
	"auth/service"
)

// @title Auth API
// @version 1.0
// @description This is sample auth server, made as dist-comp homework.

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @contact.name Arthur Petukhovsky
// @contact.url https://t.me/petuhovskiy
// @contact.email petuhovskiy@yandex.ru

// @host localhost:8081
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

	// repositories
	usersRepo := psql.NewUsers(conn)
	refreshRepo := psql.NewRefreshTokens(conn)
	confirmRepo := psql.NewConfirm(conn)

	// applying migrations
	err = usersRepo.Migrate()
	if err != nil {
		log.Panic("failed to migrate users repo, err=", err)
	}
	err = refreshRepo.Migrate()
	if err != nil {
		log.Panic("failed to migrate refresh tokens repo, err=", err)
	}
	err = confirmRepo.Migrate()
	if err != nil {
		log.Panic("failed to migrate confirm repo, err=", err)
	}

	// generators
	jwtIDGenerator := service.NewGenerator(conf.TokenLength)
	refreshGenerator := service.NewGenerator(conf.TokenLength)
	confirmGenerator := service.NewGenerator(conf.TokenLength)

	// clients
	notificator := service.NewNotificator(conf.AmqpURL, conf.QueueSms, conf.QueueEmail)

	// initializing services
	jwtSigner, err := jwt.NewHS256([]byte(conf.JWTSecret))
	if err != nil {
		log.Panic("failed to init jwt signer, err=", err)
	}
	jwtBuilder := jwt.NewTokenBuilder(jwtSigner)
	jwtService := service.NewJWT(jwtSigner, jwtBuilder, jwtIDGenerator)
	authService := service.NewAuth(
		refreshGenerator,
		refreshRepo,
		usersRepo,
		confirmGenerator,
		confirmRepo,
		notificator,
		jwtService,
		conf,
	)

	// grpc
	serv := grpcserv.New(
		conf.Grpc,
		logrus.StandardLogger(),
		&grpcserv.Services{
			Auth: grpcserv.NewAuth(authService),
		},
	)
	go func() {
		log.Fatal(serv.Listen())
	}()

	// clients gRPC
	gconn, err := grpc.Dial(
		"auth:9090", // TODO: another way
		grpc.WithInsecure(),
	)
	if err != nil {
		logrus.Fatal(err)
	}
	authCli := pb.NewAuthClient(gconn)

	// initializing handlers
	authV1 := handlers.NewAuthV1(authService)
	handler := routers.CreateRouter(authV1, authCli)

	log.Println("Serving at http://localhost" + conf.BindAddr)
	err = http.ListenAndServe(conf.BindAddr, handler)
	log.Println(err)
}
