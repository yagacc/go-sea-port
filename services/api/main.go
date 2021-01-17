package main

import (
	"context"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"github.com/yagacc/go-sea-port/domain/domain"
	pb "github.com/yagacc/go-sea-port/domain/repository/v1"
	"github.com/yagacc/go-sea-port/pkg/grpc"
	"github.com/yagacc/go-sea-port/pkg/interrupt"
	"github.com/yagacc/go-sea-port/pkg/json"
	"github.com/yagacc/go-sea-port/pkg/logger"
	"strconv"
)

type Config struct {
	GrpcPort       int    `default:"3000" split_words:"true"`
	HttpPort       int    `default:"8000" split_words:"true"`
	PortDataSource string `default:".localhost/spec/ports.json" split_words:"true"`
	RepoHost       string `default:"localhost" split_words:"true"`
	RepoPort       int    `default:"3000" split_words:"true"`
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

	//repository client
	repoAddr := config.RepoHost + ":" + strconv.Itoa(config.RepoPort)
	repoClient := &RepositoryClient{}
	repoClient.connect(context.Background(), repoAddr)

	//process json
	log.Infof("Processing Port Data [filename=%s]", config.PortDataSource)
	jsonReader := &json.BufferedJsonReader{BufferSize: 128}
	createOrUpdate := func(key string, port *domain.Port) error {
		port.Id = key
		_, err := repoClient.PortRepositoryClient.Save(context.Background(), &pb.SaveRequest{
			Port: port,
		})
		if err != nil {
			log.Errorf("cannot save %s", key)
			return err
		}
		return nil
	}
	count, err := jsonReader.ReadFile(config.PortDataSource, createOrUpdate)
	if err != nil {
		log.Fatalf("cannot process json data", err)
	}
	log.Infof("Processed Port Data [records=%d]", count)

	//start servers
	grpcServerAddr := ":" + strconv.Itoa(config.GrpcPort)
	httpServerAddr := ":" + strconv.Itoa(config.HttpPort)
	grpcServer, httpServer := NewWithShutdownOnSigTerm(httpServerAddr, grpcServerAddr)
	go grpcServer.listenAndServe(repoClient)
	go httpServer.listenAndServe()
	log.Infof("API running [http=%s] [grpc=%s]", httpServerAddr, grpcServerAddr)

	//graceful shutdown
	interrupt.ShutdownOnSigTerm() //blocks
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
