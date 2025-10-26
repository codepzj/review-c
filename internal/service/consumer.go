package service

import (
	"context"

	v1 "review-c/api/consumer/v1"
	"review-c/internal/biz"
)

type ConsumerService struct {
	v1.UnimplementedConsumerServer
	uc *biz.ConsumerUsecase
}

func NewConsumerService(uc *biz.ConsumerUsecase) *ConsumerService {
	return &ConsumerService{uc: uc}
}

// CreateReview 用户端创建评论
func (s *ConsumerService) CreateReview(ctx context.Context, req *v1.CreateConsumerRequest) (*v1.CreateConsumerReply, error) {
	id, err := s.uc.CreateReview(ctx, &biz.ReviewCreate{
		UserId:       req.UserId,
		OrderId:      req.OrderId,
		StoreId:      req.StoreId,
		Content:      req.Content,
		PicInfo:      req.PicInfo,
		VideoInfo:    req.VideoInfo,
		Score:        req.Score,
		ServiceScore: req.ServiceScore,
		ExpressScore: req.ExpressScore,
		Anonymous:    req.Anonymous,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateConsumerReply{ReviewId: id}, nil
}
