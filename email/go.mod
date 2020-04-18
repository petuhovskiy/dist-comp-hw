module email

go 1.13

require (
	github.com/caarlos0/env/v6 v6.2.1
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/sirupsen/logrus v1.5.0
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71
	golang.org/x/sys v0.0.0-20200327173247-9dae0f8f5775 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	lib v0.0.0-00010101000000-000000000000
)

replace lib => ../lib
