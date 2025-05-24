package dto

import "quote-service/iternal/domain"

type QuoteRequest struct {
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

type QuoteResponse struct {
	Id     int    `json:"id"`
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

type QuoteDeleteResponse struct {
	Message string `json:"message"`
	Id      string `json:"id"`
}

func ToResponse(q domain.Quote) QuoteResponse {
	return QuoteResponse{
		Id:     q.ID,
		Author: q.Author,
		Quote:  q.Text,
	}
}

func ToResponseSlice(src []domain.Quote) []QuoteResponse {
	out := make([]QuoteResponse, len(src))
	for i, v := range src {
		out[i] = ToResponse(v)
	}
	return out
}
