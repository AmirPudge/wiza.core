package domain

import (
	"context"
	"errors"
	"time"
)

var ErrNotFound = errors.New("not found")

type Client struct {
	ID         string    `json:"id"`
	IIN        string    `json:"iin"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	MiddleName *string   `json:"middle_name,omitempty"`
	BirthDate  time.Time `json:"birth_date"`
	Phone      *string   `json:"phone,omitempty"`
	Email      *string   `json:"email,omitempty"`
	Status     string    `json:"status"`
	Accounts   []Account `json:"accounts"`
}

type Account struct {
	ID            string    `json:"id"`
	ClientID      string    `json:"client_id"`
	AccountNumber string    `json:"account_number"`
	Currency      string    `json:"currency"`
	Balance       float64   `json:"balance"`
	Type          string    `json:"type"`
	Status        string    `json:"status"`
	OpenedAt      time.Time `json:"opened_at"`
	Cards         []Card    `json:"cards"`
}

type Card struct {
	ID             string    `json:"id"`
	AccountID      string    `json:"account_id"`
	MaskedPAN      string    `json:"masked_pan"`
	CardholderName string    `json:"cardholder_name"`
	ExpiryMonth    int       `json:"expiry_month"`
	ExpiryYear     int       `json:"expiry_year"`
	PaymentSystem  string    `json:"payment_system"`
	Type           string    `json:"type"`
	Status         string    `json:"status"`
	IssuedAt       time.Time `json:"issued_at"`
}

type ClientRepository interface {
	GetByIIN(ctx context.Context, iin string) (*Client, error)
}

type ClientService interface {
	GetByIIN(ctx context.Context, iin string) (*Client, error)
}
