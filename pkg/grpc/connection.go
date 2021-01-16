package grpc

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

type CloseFunc func()

type Connection struct {
	Timeout    time.Duration
	Add        string
	Connection *grpc.ClientConn
	close      []CloseFunc
}

func (conn *Connection) Connect(ctx context.Context) {
	if conn.Timeout == 0 {
		conn.Timeout = 3 * time.Second //default 3 seconds - zero will cause error
	}
	grpc, ctxCancelFunc, connCancelFunc := create(ctx, conn.Add, conn.Timeout)
	conn.close = append(conn.close, ctxCancelFunc)
	conn.close = append(conn.close, connCancelFunc)
	conn.Connection = grpc
}

func (conn *Connection) Stop() {
	for _, f := range conn.close {
		f()
	}
}

func create(ctx context.Context, connection string, timeout time.Duration) (conn *grpc.ClientConn, ctxCancelFunc func(), connectionCancelFunc func()) {
	ctxWithTimeout, ctxCancelFunc := context.WithTimeout(ctx, timeout)
	conn, connectionCancelFunc = dial(ctxWithTimeout, connection)
	return conn, ctxCancelFunc, connectionCancelFunc
}

func dial(ctx context.Context, address string) (*grpc.ClientConn, func()) {
	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure()) //this doesnt fail if not running
	if err != nil {
		log.Fatalf("cannot open connection [%s] [%s]", err.Error(), address)
	}
	return conn, func() {
		err = conn.Close()
		if err != nil {
			log.Println("cannot close connection properly", err)
		}
	}
}
