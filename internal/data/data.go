package data

import (
	"context"
	v1 "review-c/api/review/v1"

	"review-c/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewData, NewConsumerRepo, NewGrpcClient)

type Data struct {
	client v1.ReviewClient
}

func NewData(c *conf.Data, logger log.Logger, client v1.ReviewClient) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{client: client}, cleanup, nil
}

func NewGrpcClient(logger log.Logger ) v1.ReviewClient {

	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint("localhost:9000"), grpc.WithMiddleware(
		recovery.Recovery(),
	))
	if err != nil {
		log.NewHelper(logger).Error("连接review-service失败")
		return nil
	}
	return v1.NewReviewClient(conn)
}
