package utils

import (
	"fmt"
	"time"

	"github.com/GaryBrownEEngr/go_web_dev/backend/models"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type pasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

var _ models.TokenMaker = &pasetoMaker{}

func NewPasetoMaker(symmetricKey string) (*pasetoMaker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("Invalid key size: must be %d bytes", chacha20poly1305.KeySize)
	}
	ret := &pasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return ret, nil
}

func (s *pasetoMaker) Create(username string, duration time.Duration) (*models.Token, error) {
	payload, err := models.NewPlayload(username, duration)
	if err != nil {
		return nil, err
	}

	ret, err := s.paseto.Encrypt(s.symmetricKey, payload, nil)
	if err != nil {
		return nil, err
	}

	token := models.Token(ret)

	return &token, nil
}

func (s *pasetoMaker) Verify(token *models.Token) (*models.Payload, bool) {
	if token == nil {
		return nil, false
	}

	payload := &models.Payload{}

	err := s.paseto.Decrypt(string(*token), s.symmetricKey, payload, nil)
	if err != nil {
		return nil, false
	}

	ok := payload.Valid()
	if !ok {
		return nil, false
	}

	return payload, true
}
