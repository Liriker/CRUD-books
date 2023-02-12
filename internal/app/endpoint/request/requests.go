package request

type BookRequest struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	PublishDate int64  `json:"publish-date"`
}

func New() *BookRequest {
	return &BookRequest{}
}
