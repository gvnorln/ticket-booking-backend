package models

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	EventName string  `json:"event_name"`
	Category  string  `json:"category"`
	Location  string  `json:"location"`
	Price     float64 `json:"price"`
	Available int     `json:"available"`
}
