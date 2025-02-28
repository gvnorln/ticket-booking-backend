package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	EventID uint   `json:"event_id"`
	Status  string `json:"status"` // Contoh: "Pending", "Paid", "Failed"
}
