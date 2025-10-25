package data

import (
	"context"
	v1 "review-c/api/review/v1"

	"review-c/internal/conf"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"github.com/hashicorp/consul/api"
)

var ProviderSet = wire.NewSet(NewData, NewConsumerRepo, NewDiscovery, NewReviewClient)

type Data struct {
	client v1.ReviewClient
}

func NewData(c *conf.Data, logger log.Logger, client v1.ReviewClient) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{client: client}, cleanup, nil
}

func NewDiscovery(c *conf.Registry) registry.Discovery {
	cfg := api.DefaultConfig()
	cfg.Address = c.Addr
	cfg.Scheme = c.Scheme
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	dis := consul.New(client)
	return dis
}

func NewReviewClient(logger log.Logger, d registry.Discovery) v1.ReviewClient {
	conn, err := grpc.DialInsecure(context.Background(),
		grpc.WithDiscovery(d),
		grpc.WithEndpoint("discovery:///review-service"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		))
	if err != nil {
		log.NewHelper(logger).Error("连接review-service失败")
		return nil
	}
	return v1.NewReviewClient(conn)
}
