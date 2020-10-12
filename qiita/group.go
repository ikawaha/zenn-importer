package qiita

import (
	"time"
)

type Group struct {
	CreatedAt time.Time `json:"created_at"`
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Private   bool      `json:"private"`
	UpdatedAt time.Time `json:"updated_at"`
	URLName   string    `json:"url_name"`
}
