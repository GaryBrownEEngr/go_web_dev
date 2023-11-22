package api

import (
	"encoding/json"
	"net/http"

	"github.com/GaryBrownEEngr/go_web_dev/backend/models"
	"github.com/GaryBrownEEngr/go_web_dev/backend/utils"
	"github.com/GaryBrownEEngr/go_web_dev/backend/utils/stacktrs"
	"github.com/GaryBrownEEngr/go_web_dev/backend/utils/uerr"
)

func (s *Server) mathGameRead() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenPayload := utils.AuthMiddlewareGetTokenPayload(r.Context())
		if tokenPayload == nil || tokenPayload.Username == "" {
			err := uerr.UErrLog("Auth Error", http.StatusUnauthorized, stacktrs.Errorf("%#v", tokenPayload))
			utils.OutputError(w, err)
			return
		}

		data, err := mathGameRead(s.MathData, tokenPayload)
		if err != nil {
			utils.OutputError(w, err)
			return
		}

		utils.EncodeJSON(w, data, http.StatusOK)
	}
}

func mathGameRead(mathData models.MathGameStore, tokenPayload *models.Payload) (*models.MathGameData, error) {
	if tokenPayload == nil {
		return nil, uerr.UErrLogHash("Token Invalid", http.StatusUnauthorized, stacktrs.Errorf("Nil pointer"))
	}

	data, err := mathData.Read(tokenPayload.Username)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Server) mathGameWrite() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenPayload := utils.AuthMiddlewareGetTokenPayload(r.Context())
		if tokenPayload == nil || tokenPayload.Username == "" {
			err := uerr.UErrLog("Auth Error", http.StatusUnauthorized, stacktrs.Errorf("%#v", tokenPayload))
			utils.OutputError(w, err)
			return
		}

		var in models.MathGameData
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			http.Error(w, "unable to parse JSON", http.StatusBadRequest)
			return
		}

		err := mathGameWrite(s.MathData, tokenPayload, &in)
		if err != nil {
			utils.OutputError(w, err)
			return
		}

		utils.EncodeJSON(w, nil, http.StatusOK)
	}
}

func mathGameWrite(mathData models.MathGameStore, tokenPayload *models.Payload, in *models.MathGameData) error {
	if tokenPayload == nil || in == nil {
		return uerr.UErrLogHash("server error", http.StatusUnauthorized, stacktrs.Errorf("Nil pointer"))
	}

	in.Username = tokenPayload.Username

	err := mathData.Write(in)
	if err != nil {
		return err
	}

	return nil
}
