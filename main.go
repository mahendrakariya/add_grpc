package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/mahendrakariya/add/add"
	"github.com/mahendrakariya/add/config"
	"github.com/mahendrakariya/add/consul"
	"github.com/pkg/errors"

	"google.golang.org/grpc"
	"source.golabs.io/go-libs/service_commons/logger"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

func (s *server) DoSum(ctx context.Context, in *add.Numbers) (*add.Resp, error) {
	return &add.Resp{
		Sum: in.A + in.B + 10,
	}, nil
}

func start() error {
	config.Load()
	err := consul.InitializeClient()
	if err != nil {
		return errors.Wrap(err, "failed to initialize client")
	}

	logger.Get().Info("Starting server")
	err = consul.Register()
	if err != nil {
		return errors.Wrap(err, "failed to register with consul")
	}
	return nil
}

func main() {
	start()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	add.RegisterAdderServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		err = s.Serve(lis)
		if err != http.ErrServerClosed {
			e := consul.DeRegister()
			if e != nil {
				logger.Get().WithField("error", e.Error()).Error("failed to de-register from consul")
			}
			logger.Get().WithField("error", err.Error()).Fatal("failed to start http server")
		}
	}()
	<-sigChan

	err = consul.DeRegister()
	if err != nil {
		fmt.Printf("failed to de-register from consul: %v", err)
	}
	s.Stop()
}
