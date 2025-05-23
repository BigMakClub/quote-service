package handler

import (
	"encoding/json"
	"net/http"
	"quote-service/iternal/adapters/http/dto"
	"quote-service/iternal/domain"
	"quote-service/iternal/usecase"
)

type Handler struct {
	svc *usecase.QuoteService
}

func NewHandler(svc *usecase.QuoteService) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/quotes", h.quotes)
	mux.HandleFunc("/quotes/random", h.random)
	mux.HandleFunc("quotes", h.deleteByID)
}

func (h *Handler) quotes(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		var in dto.QuoteRequest
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		q := domain.Quote{
			Author: in.Author,
			Text:   in.Quote,
		}
		q, err := h.svc.Add(q)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		respond(w, dto.ToResponse(q), http.StatusOK)
	case http.MethodGet:
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

}

func (h *Handler) random(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) deleteByID(w http.ResponseWriter, r *http.Request) {}

func respond(w http.ResponseWriter, v any, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}
