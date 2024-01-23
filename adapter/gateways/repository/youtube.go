package repository

import (
	"context"

	"github.com/katabuchi/sample-video-api/domain/model"
	"github.com/katabuchi/sample-video-api/domain/repository"
	"github.com/katabuchi/sample-video-api/infra"
)

type youtubeRepository struct {
	YoutubeHandler
}

func NewYoutubeRepository() repository.IYoutubeRepository {
	return &youtubeRepository{&infra.YoutubeHandler{}}
}

func (yr *youtubeRepository) GetVideos(ctx context.Context) ([]model.Youtube, error) {
	// items := yr.SearchByKeyword(ctx, "English for child")
	// var lists = make([]model.Youtube, 0, len(items))
	// for _, item := range items {
	// 	lists = append(lists, model.Youtube{Id: item.Id.VideoId, Title: item.Snippet.Title})
	// }
	var lists = []model.Youtube{}
	lists = append(lists, model.Youtube{Id: "T2d3X0YQVjY", Title: "危ない合流　東京首都高・横浜周辺"})
	return lists, nil
}

func (yr *youtubeRepository) GetVideo(ctx context.Context) (string, error) {
	// items := yr.SearchByKeyword(ctx, "English for child")
	// if len(items) == 0 {
	// 	return "", nil
	// }
	// id := *items[0].Id
	// return fmt.Sprintf("https://youtube.com/watch?v=%s", id.VideoId), nil
	return "Sample URL", nil
}
