package models

import (
	"strings"
	"time"
)

type Transactions struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	UserID          uint      `gorm:"not null" json:"user_id"`
	Amount          float64   `gorm:"not null;default:0" json:"amount"`
	Currency        string    `gorm:"not null" json:"currency"`
	TransactionType string    `json:"type"`
	Category        string    `json:"category"`
	TransactionDate time.Time `json"json:"date"`
	Description     string    `json:"description"`
	Commission      float64   `gorm:"not null;default:0" json:"commission"`
}

func IsTransactionType(s string) bool {
	switch strings.TrimSpace(s) {
	case "income", "expense", "transfer":
		return true

	}
	return false

}

type Commission struct {
	TransactionID string  `json:"transaction_id"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	Type          string  `json:"type"`
	Commission    float64 `json:"commission"`
	Date          string  `json:"date"`
	Description   string  `json:"description"`
}

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}

type Transaction struct {
	ID          uint                  `gorm:"primaryKey" json:"id"`
	UserID      uint                  `json:"user_id"`
	Amount      float64               `json:"amount"`
	Currency    string                `json:"currency"`
	Type        string                `json:"type"`
	Category    string                `json:"category"`
	Date        time.Time             `json:"date"`
	Description string                `json:"description"`
	Converted   *ConvertedTransaction `json:"converted,omitempty"`
}

type ConvertedTransaction struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

func IsTransactionType(s string) bool {
	switch strings.TrimSpace(s) {
	case "income", "expense", "transfer":
		return true

	}
	return false

}
