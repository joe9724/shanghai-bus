package models

type Joke struct {
	ID int64 `json:"id"`
	Content string `json:"content"`
	CreateTime string `json:"create_time"`
}

