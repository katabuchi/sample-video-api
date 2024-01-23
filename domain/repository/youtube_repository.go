package repository

import (
	"context"

	"github.com/katabuchi/sample-video-api/domain/model"
)

type IYoutubeRepository interface {
	GetVideos(ctx context.Context) ([]model.Youtube, error)
	GetVideo(ctx context.Context) (string, error)
}
