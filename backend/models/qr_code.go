package models

import "time"

type QRCode struct {
	UserID       string    `json:"user_id" db:"user_id"`
	RandomString string    `json:"random_string" db:"random_string"`
	Status       int       `json:"status" db:"status"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type GenerateRequest struct {
	UserID string `json:"user_id" validate:"required"`
}

type VerifyRequest struct {
	QRData string `json:"qr_data" validate:"required"`
}

type VerifyResponse struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message"`
}
