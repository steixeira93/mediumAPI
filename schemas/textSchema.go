package schemas

import (
	"gorm.io/gorm"
)

type Text struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

type TextResponse {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}