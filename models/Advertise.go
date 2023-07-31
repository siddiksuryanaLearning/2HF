package models

import (
	"time"
)

type (
	// User
	Advertise struct {
		ID         uint      `json:"id" gorm:"primary_key"`
		UserID     uint      `json:"user_id"`
		VocationID uint      `json:"vocation_id"`
		Name       string    `json:"name"`
		Duration   string    `json:"duration"`
		Price      string    `json:"price"`
		Content    string    `json:"content"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
		Vocation   Vocation  `json:"vocation"`
		// AgeRatingCategory   AgeRatingCategory `json:"-"`
	}
)
