package domain

type Quote struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Text   string `json:"quote"`
}

type QuoteRepository interface {
	AddQuote(q Quote) (Quote, error)
	GetAllQuotes() ([]Quote, error)
	GetRandomQuote() (*Quote, error)
	FilterByAuthor(author string) ([]Quote, error)
	DeleteQuote(id int) error
}
