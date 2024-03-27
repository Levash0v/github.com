package models

import (
	"strings"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserResponse struct {
	Greeting string `json:"greeting"`
}

type ItemResponse struct {
	Item string `json:"id"`
	Ok   string `json:"status"`
}

type Transaction struct {
	ID          string  `json:"id"`
	Amount      float32 `json:"amount"`      
	Currency    string  `json:"currency"`    
	Types       string  `json:"type"`        
	Category    string  `json:"category"`    
	Dates       string  `json:"date"`        
	Description string  `json:"description"`

}
type TransactionList struct {
	Item []Transaction `json:"item"`
	Ok   bool          `json:"Ok"`
}

func IsTransactionType(s string) bool {
	switch strings.TrimSpace(s) {
	case "income", "expense", "transfer":
		return true

	}
	return false

}
