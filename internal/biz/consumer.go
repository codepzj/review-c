package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// C端创建评论
type ReviewCreate struct {
	UserId       int64
	OrderId      int64
	StoreId      int64
	Content      string
	PicInfo      string
	VideoInfo    string
	Score        int32
	ServiceScore int32
	ExpressScore int32
	Anonymous    int32
}

type ConsumerRepo interface {
	SaveReview(context.Context, *ReviewCreate) (int64, error)
}

type ConsumerUsecase struct {
	repo ConsumerRepo
	log  *log.Helper
}

func NewConsumerUsecase(repo ConsumerRepo, logger log.Logger) *ConsumerUsecase {
	return &ConsumerUsecase{repo: repo, log: log.NewHelper(logger)}
}

// 创建评论
func (uc *ConsumerUsecase) CreateReview(ctx context.Context, c *ReviewCreate) (int64, error) {
	uc.log.WithContext(ctx).Infof("CreateReview: %v", c)
	return uc.repo.SaveReview(ctx, c)
}
