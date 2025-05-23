package dto

type QuoteRequest struct {
	Author string `json:"author"`
	Quote  string `json:"quote"`
}
