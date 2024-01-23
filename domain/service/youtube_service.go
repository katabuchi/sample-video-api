package service

import (
	"context"

	"github.com/katabuchi/sample-video-api/domain/model"
	"github.com/katabuchi/sample-video-api/domain/repository"
)

type IYoutubeService interface {
	FindVideos(ctx context.Context) ([]model.Youtube, error)
	FindVideo(ctx context.Context) (string, error)
}

type youtubeService struct {
	repo repository.IYoutubeRepository
}

func NewYoutubeService(yr repository.IYoutubeRepository) IYoutubeService {
	return &youtubeService{repo: yr}
}

func (ys *youtubeService) FindVideos(ctx context.Context) ([]model.Youtube, error) {
	return ys.repo.GetVideos(ctx)
}

func (ys *youtubeService) FindVideo(ctx context.Context) (string, error) {
	return ys.repo.GetVideo(ctx)
}
