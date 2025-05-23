package dto

type QuoteRequest struct {
	ID     string `json:"id"`
	Author string `json:"author"`
	Quote  string `json:"quote"`
}
