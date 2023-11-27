package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/GaryBrownEEngr/go_web_dev/backend/models"
	"github.com/GaryBrownEEngr/go_web_dev/backend/utils/stacktrs"
	"github.com/GaryBrownEEngr/go_web_dev/backend/utils/uerr"
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

func (s *pasetoMaker) Verify(token *models.Token) (*models.Payload, error) {
	if token == nil {
		return nil, uerr.UErrLogHash("Token Verify Error", http.StatusInternalServerError, stacktrs.Errorf("Token is nil"))
	}

	payload := &models.Payload{}

	err := s.paseto.Decrypt(string(*token), s.symmetricKey, payload, nil)
	if err != nil {
		return nil, uerr.UErrLogHash("Token format invalid", http.StatusInternalServerError, fmt.Errorf("%#v", *token))
	}

	now := time.Now()
	if now.After(payload.ExpiredAt) {
		return nil, uerr.UErrLog("Token expired", http.StatusUnauthorized, fmt.Errorf(payload.Username))
	}

	// Make sure the token was created in the past, with 5 seconds of wiggle room.
	if now.Add(time.Second * 5).Before(payload.IssuedAt) {
		return nil, uerr.UErrLogHash("Token being used before created time", http.StatusUnauthorized, fmt.Errorf(payload.Username))
	}

	return payload, nil
}
