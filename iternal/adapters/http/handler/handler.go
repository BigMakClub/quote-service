package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"quote-service/iternal/adapters/http/dto"
	"quote-service/iternal/domain"
	"quote-service/iternal/usecase"
	"strconv"
	"strings"
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
		author := r.URL.Query().Get("author")

		var (
			list []domain.Quote
			err  error
		)

		if author == "" {
			list, err = h.svc.GetAllQuotes()
		} else {
			list, err = h.svc.FilterByAuthor(domain.Quote{Author: author})
		}

		if err != nil {
			if errors.Is(err, usecase.ErrQuoteNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		respond(w, dto.ToResponseSlice(list), http.StatusOK)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

}

func (h *Handler) random(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	q, err := h.svc.GetRandomQuote()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if q == nil {
		http.Error(w, "no quotes", http.StatusNotFound)
		return
	}

	respond(w, dto.ToResponse(*q), http.StatusOK)
}

func (h *Handler) deleteByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/quotes/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.svc.DeleteQuote(id); err != nil {
		if errors.Is(err, usecase.ErrQuoteNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func respond(w http.ResponseWriter, v any, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}
