package utils

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/GaryBrownEEngr/go_web_dev/backend/models"
	"github.com/GaryBrownEEngr/go_web_dev/backend/utils/uerr"
)

type AuthMiddleware struct {
	handler    http.Handler
	tokenMaker models.TokenMaker
}

type ContextKey string

const ContextAuthTokenPayloadKey ContextKey = "auth_token_payload"

func (s *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	tokenParts := strings.SplitN(token, " ", 2)
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		err := uerr.UErrLog("Authorization Bearer header required", http.StatusUnauthorized, fmt.Errorf(token))
		OutputError(w, err)
		return
	}
	token = tokenParts[1]
	tokenT := models.Token(token)
	tokenPayload, err := s.tokenMaker.Verify(&tokenT)
	if err != nil {
		OutputError(w, err)
		return
	}

	// create a new request context containing the authenticated user
	ctxWithUser := context.WithValue(r.Context(), ContextAuthTokenPayloadKey, tokenPayload)
	// create a new request using that new context
	rWithUser := r.WithContext(ctxWithUser)
	// call the real handler, passing the new request
	s.handler.ServeHTTP(w, rWithUser)
}

func EnsureAuth(tokenMaker models.TokenMaker, handlerToWrap http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		handler:    handlerToWrap,
		tokenMaker: tokenMaker,
	}
}

// getUser returns an instance of User,
// if set, from the given context
func AuthMiddlewareGetTokenPayload(ctx context.Context) *models.Payload {
	tokenPayload, ok := ctx.Value(ContextAuthTokenPayloadKey).(*models.Payload)
	if !ok {
		return nil
	}
	return tokenPayload
}
