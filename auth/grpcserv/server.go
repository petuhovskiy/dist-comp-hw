package grpcserv

import (
	"auth/config"
	"context"
	"fmt"
	"lib/pb"
	"net"
	"runtime/debug"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

const (
	alivePeriod = 10
)

type Services struct {
	Auth pb.AuthServer
}

type Server struct {
	apiServer *grpc.Server
	config    config.Grpc
}

func New(config config.Grpc, logger *log.Logger, services *Services) *Server {
	logEntry := log.NewEntry(logger)

	apiServer := grpc.NewServer(
		grpc.ConnectionTimeout(time.Second*time.Duration(config.Timeout)),
		// MiddleWare
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_ctxtags.UnaryServerInterceptor(),
				grpc_prometheus.UnaryServerInterceptor,
				grpc_logrus.UnaryServerInterceptor(logEntry),
				grpc_logrus.PayloadUnaryServerInterceptor(logEntry, func(ctx context.Context, fullMethodName string, servingObject interface{}) bool {
					return true
				}),
				grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(func(i interface{}) error {
					log.
						WithField("panic_stack", string(debug.Stack())).
						Error("panic")
					return fmt.Errorf("%#v", i)
				})),
			),
		),
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_ctxtags.StreamServerInterceptor(),
				grpc_prometheus.StreamServerInterceptor,
				grpc_logrus.StreamServerInterceptor(logEntry),
				grpc_logrus.PayloadStreamServerInterceptor(logEntry, func(ctx context.Context, fullMethodName string, servingObject interface{}) bool {
					return true
				}),
				grpc_recovery.StreamServerInterceptor(),
			),
		),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: time.Second * alivePeriod,
			Time:              time.Second * alivePeriod,
			Timeout:           time.Second * alivePeriod,
		}),
	)

	s := &Server{
		apiServer: apiServer,
		config:    config,
	}

	s.registerGRPC(services)
	return s
}

func (s *Server) Listen() error {
	listener, err := net.Listen("tcp", s.config.Bind)
	if err != nil {
		log.WithError(err).Error("gRPC listen")
		return err
	}

	log.Info("gRPC server started")
	defer log.Info("gRPC server exited")
	if err := s.apiServer.Serve(listener); err != nil {
		log.WithError(err).Error("gRPC serve")
		return err
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.apiServer.GracefulStop()
	return nil
}

func (s *Server) registerGRPC(ss *Services) {
	pb.RegisterAuthServer(s.apiServer, ss.Auth)
	reflection.Register(s.apiServer)
}
