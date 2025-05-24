package main

import (
	"log"
	"net/http"
	"quote-service/config"
	"quote-service/iternal/adapters/http/handler"
	"quote-service/iternal/adapters/storage"
	"quote-service/iternal/usecase"
)

func main() {
	cfg, err := config.LoadConfig("./config/config.json")
	if err != nil {
		log.Fatal(err)
	}
	repo, err := storage.NewStorage(cfg.StoragePath)
	if err != nil {
		log.Fatal(err)
	}
	svc := usecase.NewQuoteService(repo)
	h := handler.NewHandler(svc)
	mux := http.NewServeMux()
	h.Register(mux)
	addr := ":8080"
	log.Printf("Quote-Service listening at http://localhost%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
