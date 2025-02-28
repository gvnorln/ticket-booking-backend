package models

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Location string  `json:"location"`
	Price    float64 `json:"price"`
	Seats    int     `json:"seats"`
	Action   string  `json:"action"`
}
