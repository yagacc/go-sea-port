package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"github.com/yagacc/go-sea-port/pkg/grpc"
	"github.com/yagacc/go-sea-port/pkg/interrupt"
	"github.com/yagacc/go-sea-port/pkg/logger"
	"os"
	"strconv"
)

type Config struct {
	GrpcPort int `default:"3000" split_words:"true"`
	HttpPort int `default:"8888" split_words:"true"`
}

var (
	logrusLog = logrus.New()
	log       = logger.NewLogrusLogger(logrusLog)
)

func main() {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		fatal("cannot process env vars", err)
	}
	grpcServerAddr := ":" + strconv.Itoa(config.GrpcPort)
	httpServerAddr := ":" + strconv.Itoa(config.HttpPort)
	grpcServer, httpServer := NewWithShutdownOnSigTerm(httpServerAddr, grpcServerAddr)
	go grpcServer.listenAndServe()
	go httpServer.listenAndServe()
	log.Infof("API running [http=%s] [grpc=%s]", httpServerAddr, grpcServerAddr)

	interrupt.ShutdownOnSigTerm()
	grpcServer.server.Stop()
	httpServer.stop()
	log.Infof("Shut down")
}

func NewWithShutdownOnSigTerm(httpAddr, grpcAddr string) (grpcSrv *GrpcServer, httpSrv *HttpServer) {
	grpcSrv = &GrpcServer{
		server: grpc.NewServer(grpcAddr),
	}
	httpSrv = &HttpServer{
		address:        httpAddr,
		grpcServerAddr: grpcAddr,
		interrupt:      interrupt.ShutdownOnSigTerm,
	}
	return
}

func fatal(fmt string, v interface{}) {
	log.Errorf(fmt, v)
	os.Exit(1)
}
