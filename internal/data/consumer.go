package data

import (
	"context"

	v1 "review-c/api/review/v1"
	"review-c/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type consumerRepo struct {
	data *Data
	log  *log.Helper
}

// NewConsumerRepo .
func NewConsumerRepo(data *Data, logger log.Logger) biz.ConsumerRepo {
	return &consumerRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *consumerRepo) SaveReview(ctx context.Context, g *biz.ReviewCreate) (int64, error) {
	review, err := r.data.client.CreateReview(ctx, &v1.CreateReviewRequest{
		UserId:       g.UserId,
		OrderId:      g.OrderId,
		StoreId:      g.StoreId,
		Content:      g.Content,
		PicInfo:      g.PicInfo,
		VideoInfo:    g.VideoInfo,
		Score:        g.Score,
		ServiceScore: g.ServiceScore,
		ExpressScore: g.ExpressScore,
		Anonymous:    g.Anonymous,
	})
	if err != nil {
		return 0, err
	}
	return review.ReviewId, nil
}
