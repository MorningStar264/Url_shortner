package model

import "time"

type Login struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
}
type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Link struct {
	ID             int       `json:"id"`
	ShortCode      string    `json:"short_code"`
	LongURL        string    `json:"long_url"`
	ClicksCount    int       `json:"clicks_count"`
	LastAccessedAt time.Time `json:"last_accessed_at"`
	CreatedBy      int       `json:"created_by"`
	CreatedAt      time.Time `json:"created_at"`
	ExpiresAt      time.Time `json:"expires_at"`
	IsActive       bool      `json:"is_active"`
}
