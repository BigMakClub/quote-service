package storage

import (
	"errors"
	"math/rand"
	"quote-service/iternal/domain"
	"slices"
	"sync"
)

type Storage struct {
	data   []domain.Quote
	nextID int
	mut    sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		data:   make([]domain.Quote, 0),
		nextID: 1,
	}
}

func (s *Storage) AddQuote(quote domain.Quote) (domain.Quote, error) {
	s.mut.Lock()
	defer s.mut.Unlock()
	quote.ID = s.nextID
	s.nextID++
	s.data = append(s.data, quote)
	return quote, nil
}

func (s *Storage) GetAllQuotes() ([]domain.Quote, error) {
	s.mut.RLock()
	defer s.mut.RUnlock()
	out := make([]domain.Quote, len(s.data))
	copy(out, s.data)
	return out, nil
}

func (s *Storage) GetRandomQuote() (*domain.Quote, error) {
	s.mut.RLock()
	defer s.mut.RUnlock()
	if len(s.data) == 0 {
		return nil, errors.New("цитата не найдена")
	}
	out := s.data[rand.Intn(len(s.data))]
	return &out, nil
}

func (s *Storage) FilterByAuthor(author string) ([]domain.Quote, error) {
	s.mut.RLock()
	defer s.mut.RUnlock()
	out := make([]domain.Quote, 0)
	for _, quote := range s.data {
		if quote.Author == author {
			out = append(out, quote)
		}
	}
	if len(out) == 0 {
		return nil, errors.New("цитаты автора не найдены")
	}
	return out, nil
}

func (s *Storage) DeleteQuote(id int) error {
	s.mut.Lock()
	defer s.mut.Unlock()

	for i, quote := range s.data {
		if quote.ID == id {
			//s.data = append(s.data[:i], s.data[i+1:]...)
			slices.Delete(s.data, i, i)
			return nil
		}
	}

	return errors.New("нет такого id")

}
