package models

import (
	"time"
)

type (
	// User
	Payment struct {
		ID          uint      `json:"id" gorm:"primary_key"`
		UserID      uint      `json:"user_id"`
		VocationID  *uint     `json:"vocation_id"`
		AdvertiseID *uint     `json:"advertise_id"`
		Amount      float64   `json:"amount"`
		Currency    string    `json:"currency"`
		DueDate     string    `json:"due_date"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		Advertise   Advertise `json:"advertise"`
		Vocation    Vocation  `json:"vocation"`

		// AgeRatingCategory   AgeRatingCategory `json:"-"`
	}
)
