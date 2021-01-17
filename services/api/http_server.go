package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	api "github.com/yagacc/go-sea-port/domain/api/v1"
	"github.com/yagacc/go-sea-port/pkg/interrupt"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

type HttpServer struct {
	httpServer              *http.Server
	address, grpcServerAddr string
	interrupt               interrupt.Interrupt
}

func (s *HttpServer) listenAndServe() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: false}))
	err := initRestApi(mux, ctx, s.grpcServerAddr)
	if err != nil {
		log.Fatalf("failed to start http server: %v", err)
	}

	server := http.NewServeMux()
	server.Handle("/", mux)

	s.httpServer = &http.Server{
		Addr:    s.address,
		Handler: server,
	}
	err = s.httpServer.ListenAndServe()
	if err == nil {
		log.Fatalf("error with http server: %v", err)
	}
}

func (s *HttpServer) stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("cannot shutdown http server %v", err)
	}
	cancel()
}

func initRestApi(mux *runtime.ServeMux, ctx context.Context, localGrpcServer string) error {
	mux.GetForwardResponseOptions()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := api.RegisterPortApiHandlerFromEndpoint(ctx, mux, localGrpcServer, opts)
	if err != nil {
		log.Fatalf("failed to start http server: %v", err)
	}
	return err
}
