package models

import (
	"time"
)

type (
	// User
	Vocation struct {
		ID          uint      `json:"id" gorm:"primary_key"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Comment     string    `json:"comment"`
		Rating      string    `json:"rating"`
		Phone       string    `json:"phone"`
		Country     string    `json:"country"`
		City        string    `json:"city"`
		Address     string    `json:"address"`
		Image       string    `json:"image"`
		Price       string    `json:"price"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		UserID      uint      `json:"user_id"`

		// AgeRatingCategory   AgeRatingCategory `json:"-"`
	}
)
