package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"github.com/yagacc/go-sea-port/pkg/grpc"
	"github.com/yagacc/go-sea-port/pkg/interrupt"
	"github.com/yagacc/go-sea-port/pkg/logger"
	"strconv"
)

type Config struct {
	GrpcPort    int    `default:"3000" split_words:"true"`
}

var (
	logrusLog = logrus.New()
	log       = logger.NewLogrusLogger(logrusLog)
)

func main() {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		log.Fatalf("cannot process env vars", err)
	}

	grpcServerAddr := ":" + strconv.Itoa(config.GrpcPort)
	grpcServer := &GrpcServer{
		server: grpc.NewServer(grpcServerAddr),
	}

	go grpcServer.listenAndServe(NewInMemPortStorage())

	log.Infof("Repository running [grpc=%s]", grpcServerAddr)

	//graceful shutdown
	interrupt.ShutdownOnSigTerm() //blocks
	grpcServer.server.Stop()

	log.Infof("Shut down")
}
