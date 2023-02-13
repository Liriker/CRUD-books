package book

import "encoding/json"

type Book struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	PublishDate int64  `json:"publish-date"`
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

func (b *Book) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, b)
	return err
}
