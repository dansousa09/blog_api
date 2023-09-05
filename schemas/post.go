package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	*gorm.Model
	ID      uint
	Title   string
	Content string
	Author  string
}

type PostResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}
