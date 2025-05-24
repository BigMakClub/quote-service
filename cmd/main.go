package main

import (
	"log"
	"net/http"
	"quote-service/iternal/adapters/http/handler"
	"quote-service/iternal/adapters/storage"
	"quote-service/iternal/usecase"
)

func main() {
	repo, _ := storage.NewStorage("./db/quotes.json")
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
