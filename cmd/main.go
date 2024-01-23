package main

import (
	"net/http"

	"github.com/katabuchi/sample-video-api/adapter/controller"
)

var tc = controller.NewYoutubeController()

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/videos", tc.GetList)
	http.HandleFunc("/video/", tc.GetVideo)
	server.ListenAndServe()
}
