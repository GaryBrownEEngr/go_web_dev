package utils

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
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
	var userErr *UserErr
	if errors.As(err, &userErr) {
		errorMsg, code := userErr.UserMsgAndCode()
		http.Error(w, errorMsg, code)
		if userErr.shouldLog {
			log.Println(UnwrapAllErrorsForLog(err))
		}
		return
	}

	errorMsg := "Error: " + HashError(err)
	http.Error(w, errorMsg, code)
}
