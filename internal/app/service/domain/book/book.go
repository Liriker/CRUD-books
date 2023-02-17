package book

import "time"

type Book struct {
	Id          int       `json:"id,omitempty"`
	Name        string    `json:"name"`
	Author      string    `json:"author"`
	PublishDate time.Time `json:"publish-date"`
}

func New() *Book {
	return &Book{}
}

func (b *Book) ID() int {
	return b.Id
}

func (b *Book) UpdateBook(newBook Book) {
	if newBook.Name != "" {
		b.Name = newBook.Name
	}
	if newBook.Author != "" {
		b.Author = newBook.Author
	}
	if !newBook.PublishDate.IsZero() {
		b.PublishDate = newBook.PublishDate
	}
}
