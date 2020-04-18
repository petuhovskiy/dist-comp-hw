module app

go 1.13

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/go-chi/chi v4.0.3+incompatible
	github.com/go-chi/render v1.0.1
	github.com/go-openapi/spec v0.19.6 // indirect
	github.com/go-openapi/swag v0.19.7 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/jackc/pgx/v4 v4.4.1
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/sirupsen/logrus v1.5.0
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71
	github.com/swaggo/http-swagger v0.0.0-20200103000832-0e9263c4b516
	github.com/swaggo/swag v1.6.5
	golang.org/x/net v0.0.0-20200222125558-5a598a2470a0 // indirect
	golang.org/x/tools v0.0.0-20200221224223-e1da425f72fd // indirect
	google.golang.org/grpc v1.28.1
	gopkg.in/yaml.v2 v2.2.8 // indirect
	lib v0.0.0-00010101000000-000000000000
)

replace lib => ../lib
