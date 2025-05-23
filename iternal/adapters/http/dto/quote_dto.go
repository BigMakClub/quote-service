package dto

import "quote-service/iternal/domain"

type QuoteRequest struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

type QuoteResponse struct {
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

func ToResponse(q domain.Quote) QuoteResponse {
	return QuoteResponse{
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
