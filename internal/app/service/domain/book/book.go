package book

import "time"

type Book struct {
	Id          int       `json:"id,omitempty"`
	Name        string    `json:"name"`
	Author      string    `json:"author"`
	PublishDate time.Time `json:"publish-date"`
}

//func New(id int, name, author string, publishDate int64) *Book {
//	return &Book{
//		Id:          id,
//		Name:        name,
//		Author:      author,
//		PublishDate: publishDate,
//	}
//}

func New() *Book {
	return &Book{}
}

func (b *Book) ID() int {
	return b.Id
}
