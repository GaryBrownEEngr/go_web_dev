package utils

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/GaryBrownEEngr/go_web_dev/backend/utils/uerr"
)

func EncodeJSON(w http.ResponseWriter, thing interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(thing)
	if err != nil {
		log.Println(err)
	}
}

// This looks for a wrapped userErr and uses its message and code if possible.
// Otherwise the error is replaced with a hash and the code is StatusInternalServerError
func OutputError(w http.ResponseWriter, err error) {
	OutputErrorAndCode(w, err, http.StatusInternalServerError)
}

// This looks for a wrapped userErr and uses its message and code if possible.
// Otherwise the error is replaced with a hash
func OutputErrorAndCode(w http.ResponseWriter, err error, code int) {
	var userErr *uerr.UserErrorData
	if errors.As(err, &userErr) {
		errorMsg, code := userErr.UserMsgAndCode()
		http.Error(w, errorMsg, code)
		if userErr.ShouldLog() {
			log.Println(uerr.UnwrapAllErrorsForLog(err))
		}
		return
	}

	errorMsg := "Error: " + uerr.HashError(err)
	http.Error(w, errorMsg, code)
}
