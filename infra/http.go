package infra

import (
	"encoding/json"
	"log"
	"net/http"
)

func SetCors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "ControlType")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,UPDATE,OPTIONS")
	w.Header().Set("Content-Type", "apllication/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode("jphe"); err != nil {
		log.Println(err)
	}
}
