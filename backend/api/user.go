package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/GaryBrownEEngr/go_web_dev/backend/models"
	"github.com/GaryBrownEEngr/go_web_dev/backend/sessionuser"
	"github.com/GaryBrownEEngr/go_web_dev/backend/utils"
	"github.com/GaryBrownEEngr/go_web_dev/backend/utils/stacktrs"
)

func (s *Server) userCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type inT struct {
			UserName string `json:"username"`
			Password string `json:"password"`
		}

		var in inT
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			http.Error(w, "unable to parse JSON", http.StatusBadRequest)
			return
		}

		user, err := userCreate(s.users, in.UserName, in.Password)
		if err != nil {
			utils.OutputError(w, err)
			return
		}

		utils.EncodeJSON(w, user, http.StatusOK)
	}
}

func userCreate(users models.UserStore, username, password string) (*models.User, error) {
	user, err := users.CreateUser(username, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Server) userLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type inT struct {
			UserName string `json:"username"`
			Password string `json:"password"`
		}

		var in inT
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			http.Error(w, "unable to parse JSON", http.StatusBadRequest)
			return
		}

		user, sessionToken, err := userLogin(s.users, s.tokenMaker, in.UserName, in.Password)
		if err != nil {
			utils.OutputError(w, err)
			return
		}

		type outT struct {
			User         *models.User  `json:"user"`
			SessionToken *models.Token `json:"session_token"`
		}

		out := outT{
			User:         user,
			SessionToken: sessionToken,
		}
		utils.EncodeJSON(w, out, http.StatusOK)
	}
}

func userLogin(users models.UserStore, tokenMaker models.TokenMaker, username, password string) (*models.User, *models.Token, error) {
	user, err := users.VerifyPassword(username, password)
	if err != nil {
		return nil, nil, err
	}

	sessionToken, err := sessionuser.CreateSession(tokenMaker, user)
	if err != nil {
		return nil, nil, err
	}
	return user, sessionToken, nil
}

func (s *Server) userDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		tokenParts := strings.SplitN(token, " ", 2)
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			err := utils.NewUserErrLog("Authorization Bearer header required", http.StatusUnauthorized, stacktrs.Errorf(token))
			utils.OutputError(w, err)
			return
		}
		token = tokenParts[1]

		tokenT := models.Token(token)
		err := userDelete(s.users, s.tokenMaker, &tokenT)
		if err != nil {
			utils.OutputError(w, err)
			return
		}

		utils.EncodeJSON(w, nil, http.StatusOK)
	}
}

func userDelete(users models.UserStore, tokenMaker models.TokenMaker, token *models.Token) error {
	if token == nil {
		return utils.NewUserErrLogHash("Token Invalid", http.StatusUnauthorized, stacktrs.Errorf("Nil pointer"))
	}

	payload, ok := tokenMaker.Verify(token)
	if !ok {
		return utils.NewUserErrLogHash("Token Invalid", http.StatusUnauthorized, stacktrs.Errorf(string(*token)))
	}

	err := users.DeleteUser(payload.Username)
	if err != nil {
		return err
	}

	return nil
}
