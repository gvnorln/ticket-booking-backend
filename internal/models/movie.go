package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Duration    int     `json:"duration"`
	Genre       string  `json:"genre"`
	Rating      float64 `json:"rating"`
}
