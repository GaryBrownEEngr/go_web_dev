package models

import (
	"fmt"
	"net/http"
	"time"

	"github.com/GaryBrownEEngr/go_web_dev/backend/utils/uerr"
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

func (s *Payload) Valid() error {
	now := time.Now()
	if now.After(s.ExpiredAt) {
		return uerr.UErrLog("Token expired", http.StatusUnauthorized, fmt.Errorf(s.Username))
	}

	// Make sure the token was created in the past, with 5 seconds of wiggle room.
	if now.Add(time.Second * 5).Before(s.IssuedAt) {
		return uerr.UErrLogHash("Token being used before created time", http.StatusUnauthorized, fmt.Errorf(s.Username))
	}

	return nil
}

type TokenMaker interface {
	Create(username string, duration time.Duration) (*Token, error)
	Verify(token *Token) (*Payload, error)
}
