package usecase

import (
	"quote-service/iternal/domain"
	"testing"
)

type MockQuoteRepository struct {
	quotes []domain.Quote
}

func (m *MockQuoteRepository) AddQuote(q domain.Quote) (domain.Quote, error) {
	m.quotes = append(m.quotes, q)
	return q, nil
}

func (m *MockQuoteRepository) GetAllQuotes() ([]domain.Quote, error) {
	return m.quotes, nil
}

func (m *MockQuoteRepository) GetRandomQuote() (*domain.Quote, error) {
	if len(m.quotes) == 0 {
		return nil, ErrQuoteNotFound
	}
	return &m.quotes[0], nil
}

func (m *MockQuoteRepository) FilterByAuthor(author string) ([]domain.Quote, error) {
	var result []domain.Quote
	for _, q := range m.quotes {
		if q.Author == author {
			result = append(result, q)
		}
	}
	return result, nil
}

func (m *MockQuoteRepository) DeleteQuote(id int) error {
	for i, q := range m.quotes {
		if q.ID == id {
			m.quotes = append(m.quotes[:i], m.quotes[i+1:]...)
			return nil
		}
	}
	return ErrQuoteNotFound
}

func TestQuoteService_Add(t *testing.T) {
	tests := []struct {
		name    string
		quote   domain.Quote
		wantErr bool
	}{
		{
			name: "valid quote",
			quote: domain.Quote{
				Author: "Test Author",
				Text:   "Test Quote",
			},
			wantErr: false,
		},
		{
			name: "empty author",
			quote: domain.Quote{
				Author: "",
				Text:   "Test Quote",
			},
			wantErr: true,
		},
		{
			name: "empty text",
			quote: domain.Quote{
				Author: "Test Author",
				Text:   "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MockQuoteRepository{}
			service := NewQuoteService(repo)
			_, err := service.Add(tt.quote)
			if (err != nil) != tt.wantErr {
				t.Errorf("QuoteService.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQuoteService_FilterByAuthor(t *testing.T) {
	repo := &MockQuoteRepository{
		quotes: []domain.Quote{
			{ID: 1, Author: "Author1", Text: "Quote1"},
			{ID: 2, Author: "Author2", Text: "Quote2"},
			{ID: 3, Author: "Author1", Text: "Quote3"},
		},
	}
	service := NewQuoteService(repo)

	tests := []struct {
		name    string
		author  string
		want    int
		wantErr bool
	}{
		{
			name:    "existing author",
			author:  "Author1",
			want:    2,
			wantErr: false,
		},
		{
			name:    "non-existing author",
			author:  "Author3",
			want:    0,
			wantErr: false,
		},
		{
			name:    "empty author",
			author:  "",
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quotes, err := service.FilterByAuthor(domain.Quote{Author: tt.author})
			if (err != nil) != tt.wantErr {
				t.Errorf("QuoteService.FilterByAuthor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(quotes) != tt.want {
				t.Errorf("QuoteService.FilterByAuthor() got %d quotes, want %d", len(quotes), tt.want)
			}
		})
	}
}

func TestQuoteService_DeleteQuote(t *testing.T) {
	repo := &MockQuoteRepository{
		quotes: []domain.Quote{
			{ID: 1, Author: "Author1", Text: "Quote1"},
			{ID: 2, Author: "Author2", Text: "Quote2"},
		},
	}
	service := NewQuoteService(repo)

	tests := []struct {
		name    string
		id      int
		wantErr bool
	}{
		{
			name:    "existing quote",
			id:      1,
			wantErr: false,
		},
		{
			name:    "non-existing quote",
			id:      3,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.DeleteQuote(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("QuoteService.DeleteQuote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
