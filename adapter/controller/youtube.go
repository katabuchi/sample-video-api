package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/katabuchi/sample-video-api/adapter/gateways/repository"
	"github.com/katabuchi/sample-video-api/domain/service"
	usecases "github.com/katabuchi/sample-video-api/usecase"
)

type YoutubeController interface {
	GetList(w http.ResponseWriter, r *http.Request)
	GetVideo(w http.ResponseWriter, r *http.Request)
}

type youtubeController struct {
	usecase usecases.IYoutubeUseCase
}

func NewYoutubeController() YoutubeController {
	repository := repository.NewYoutubeRepository()
	service := service.NewYoutubeService(repository)
	usecase := usecases.NewYoutubeUseCase(service)
	return &youtubeController{usecase}
}

func (yc *youtubeController) GetList(w http.ResponseWriter, r *http.Request) {
	context := context.Background()
	items, err := yc.usecase.GetYoutube(context)
	if err != nil {
		log.Fatalf("GetYoutube has Error: %v", err)
	}
	for _, item := range items {
		fmt.Printf("item Id : %s, Title: %s", item.Id, item.Title)
	}
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,UPDATE,OPTIONS")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(items[0]); err != nil {
		log.Println(err)
	}
}

func (yc *youtubeController) GetVideo(w http.ResponseWriter, r *http.Request) {
	context := context.Background()
	url, err := yc.usecase.GetVideo(context)
	if err != nil {
		log.Fatalf("Error : %v", err)
	}
	fmt.Printf("URL is %s", url)
}
