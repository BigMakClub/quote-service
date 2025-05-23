package usecase

import (
	"errors"
	"quote-service/iternal/domain"
)

type QuoteRepository interface {
	AddQuote(q domain.Quote) (domain.Quote, error)
	GetAllQuotes() ([]domain.Quote, error)
	GetRandomQuote() (*domain.Quote, error)
	FilterByAuthor(author string) ([]domain.Quote, error)
	DeleteQuote(id int) error
}

type QuoteService struct {
	quoteRepository QuoteRepository
}

func NewQuoteService(quoteRepository QuoteRepository) *QuoteService {
	return &QuoteService{
		quoteRepository: quoteRepository,
	}
}

func (s *QuoteService) Add(q domain.Quote) (domain.Quote, error) {
	if q.Author == "" || q.Text == "" {
		return domain.Quote{}, errors.New("invalid quote")
	}
	return s.quoteRepository.AddQuote(q)
}

func (s *QuoteService) GetAllQuotes() ([]domain.Quote, error) {
	return s.quoteRepository.GetAllQuotes()
}

func (s *QuoteService) GetRandomQuote() (*domain.Quote, error) {
	return s.quoteRepository.GetRandomQuote()
}

func (s *QuoteService) FilterByAuthor(q domain.Quote) ([]domain.Quote, error) {
	if q.Author != "" {
		return s.quoteRepository.FilterByAuthor(q.Author)
	}
	return []domain.Quote{}, errors.New("invalid quote")
}

func (s *QuoteService) DeleteQuote(id int) error {
	return s.quoteRepository.DeleteQuote(id)
}
