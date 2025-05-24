package storage

import (
	"encoding/json"
	"errors"
	"math/rand"
	"os"
	"quote-service/iternal/domain"
	"sync"
)

type Storage struct {
	file   string
	data   []domain.Quote
	nextID int
	mut    sync.RWMutex
}

func NewStorage(path string) (*Storage, error) {
	s := &Storage{
		file:   path,
		data:   make([]domain.Quote, 0),
		nextID: 1,
	}

	if path != "" {
		if err := s.load(); err != nil && !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}
	}
	return s, nil
}

func (s *Storage) AddQuote(quote domain.Quote) (domain.Quote, error) {
	s.mut.Lock()
	defer s.mut.Unlock()
	quote.ID = s.nextID
	s.nextID++
	s.data = append(s.data, quote)
	return quote, s.saveToJSON()
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
			s.data = append(s.data[:i], s.data[i+1:]...)
			return s.saveToJSON()
		}
	}

	return errors.New("нет такого id")

}

func (s *Storage) load() error {
	b, err := os.ReadFile(s.file)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, &s.data); err != nil {
		return err
	}

	for _, q := range s.data {
		if q.ID >= s.nextID {
			s.nextID = q.ID + 1
		}
	}
	return nil
}

func (s *Storage) saveToJSON() error {
	if s.file == "" {
		return nil
	}
	b, err := json.MarshalIndent(s.data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.file, b, 0644)
}
