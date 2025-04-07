package domain

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Name      string
	Email     string
	APIKey    string
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func generateAPIKey() string {
	b := make([]byte, 16) // isso é um slice de bytes cria um array de 16 bytes
	rand.Read(b)          // preenche o array com bytes aleatórios
	return hex.EncodeToString(b)
}

func NewAccount(name, email string) *Account {
	account := &Account{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Balance:   0,
		APIKey:    generateAPIKey(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return account
}

func (a *Account) AddBalance(amount float64) {
	a.Balance += amount
	a.UpdatedAt = time.Now()
}
