package rpc

import (
	"context"
	"sync"
	"time"

	v1 "subway_intro/api/user/v1"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	srcgrpc "google.golang.org/grpc"
)

type RpcClient struct {
	ctx     context.Context
	connMap sync.Map
}

func NewClient(ctx context.Context) *RpcClient {
	client := &RpcClient{
		ctx: ctx,
	}

	return client
}

func (r *RpcClient) newConn(target string) error {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///"+target),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
		),
		grpc.WithTimeout(3*time.Second),
		// for tracing remote ip recording
		grpc.WithOptions(srcgrpc.WithStatsHandler(&tracing.ClientHandler{})),
	)
	if err != nil {
		return err
	}
	r.connMap.Store(target, conn)
	go func(c *srcgrpc.ClientConn) {
		<-r.ctx.Done()
		c.Close()
	}(conn)

	return nil
}

func (r *RpcClient) getConn(target string) (*srcgrpc.ClientConn, error) {
	if _, ok := r.connMap.Load(target); !ok {
		err := r.newConn(target)
		if err != nil {
			return nil, err
		}
	}
	conn, _ := r.connMap.Load(target)

	return conn.(*srcgrpc.ClientConn), nil
}

func (r *RpcClient) UserServiceClient() (v1.UserServiceClient, error) {
	client, err := r.getConn("lemon")
	if err != nil {
		return nil, err
	}

	return v1.NewUserServiceClient(client), nil
}
