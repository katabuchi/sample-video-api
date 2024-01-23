package repository

import (
	"context"

	"google.golang.org/api/youtube/v3"
)

type YoutubeHandler interface {
	SearchByKeyword(ctx context.Context, query string) []*youtube.SearchResult
}
