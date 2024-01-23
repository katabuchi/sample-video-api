package infra

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type YoutubeHandler struct{}

func (yr *YoutubeHandler) SearchByKeyword(ctx context.Context, query string) []*youtube.SearchResult {
	apiKey := getApiKey()
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	handleError(err, "Error creating YouTube client")
	call := service.Search.List([]string{"snippet"}). // 検索
								Q(query).           // 検索クエリ
								Type("video").      // 検索対象の種類
								Order("viewCount"). // 検索結果を閲覧数でソート
								MaxResults(5)       // 取得する検索結果の動画数の最大
	response, err := call.Do()

	handleError(err, "")

	return response.Items
}

func getApiKey() string {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Unable to read env file: %v", err)
	}

	apiKey := os.Getenv("API_KEY")
	return apiKey
}

func handleError(err error, message string) {
	if message == "" {
		message = "Error making API call"
	}
	if err != nil {
		log.Fatalf(message+": %v", err.Error())
	}
}
