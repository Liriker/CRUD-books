package domain

import "time"

type Book struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Author       string    `json:"author"`
	PublishHouse string    `json:"publish-house"`
	PublishDate  time.Time `json:"publish-date"`
}

func NewBook(id int, name, author, publishHouse string, publishDate time.Time) *Book {
	return &Book{
		Id:           id,
		Name:         name,
		Author:       author,
		PublishHouse: publishHouse,
		PublishDate:  publishDate,
	}
}

func EmptyBook() *Book {
	return &Book{
		Id:           0,
		Name:         "",
		Author:       "",
		PublishHouse: "",
		PublishDate:  time.Time{},
	}
}

func (b *Book) GetID() int {
	return b.Id
}
