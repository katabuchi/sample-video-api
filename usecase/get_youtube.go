package usecases

import (
	"context"

	"github.com/katabuchi/sample-video-api/domain/model"
	"github.com/katabuchi/sample-video-api/domain/service"
)

type IYoutubeUseCase interface {
	GetYoutube(ctx context.Context) ([]model.Youtube, error)
	GetVideo(ctx context.Context) (string, error)
}

type youtubeUseCase struct {
	svc service.IYoutubeService
}

func NewYoutubeUseCase(ss service.IYoutubeService) IYoutubeUseCase {
	return &youtubeUseCase{
		svc: ss,
	}
}

func (ss *youtubeUseCase) GetYoutube(ctx context.Context) ([]model.Youtube, error) {
	return ss.svc.FindVideos(ctx)
}

func (ss *youtubeUseCase) GetVideo(ctx context.Context) (string, error) {
	return ss.svc.FindVideo(ctx)
}
