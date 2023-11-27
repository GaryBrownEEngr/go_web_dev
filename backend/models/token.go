package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Token string

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPlayload(username string, duration time.Duration) (*Payload, error) {
	if username == "" {
		return nil, fmt.Errorf("A username is required")
	}

	if duration < 0 {
		return nil, fmt.Errorf("The duration must be positive: %v", duration)
	}

	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	ret := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return ret, nil
}

type TokenMaker interface {
	Create(username string, duration time.Duration) (*Token, error)
	Verify(token *Token) (*Payload, error)
}
