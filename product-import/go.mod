module product-import

go 1.13

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/davecgh/go-spew v1.1.1
	github.com/go-chi/chi v4.0.3+incompatible
	github.com/go-chi/render v1.0.1
	github.com/go-openapi/spec v0.19.7 // indirect
	github.com/go-openapi/swag v0.19.8 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/sirupsen/logrus v1.5.0
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71
	github.com/stretchr/testify v1.5.1
	github.com/swaggo/http-swagger v0.0.0-20200308142732-58ac5e232fba
	github.com/swaggo/swag v1.6.5
	golang.org/x/net v0.0.0-20200320181208-1c781a10960a // indirect
	golang.org/x/sys v0.0.0-20190826190057-c7b8b68b1456 // indirect
	golang.org/x/tools v0.0.0-20200319210407-521f4a0cd458 // indirect
	google.golang.org/grpc v1.28.1
	gopkg.in/yaml.v2 v2.2.8 // indirect
	lib v0.0.0-00010101000000-000000000000
)

replace lib => ../lib
