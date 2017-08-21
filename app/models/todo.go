package models

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Done        bool   `json:"done"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
