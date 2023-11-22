package uerr

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

func UnwrapAllErrorsForLog(err error) string {
	if err == nil {
		return ""
	}

	msg := err.Error()

	for {
		err = errors.Unwrap(err)
		if err == nil {
			break
		}
		msg = msg + " :: " + err.Error()
	}

	return msg
}

func HashError(in error) string {
	return HashErrorString(in.Error())
}

func HashErrorString(in string) string {
	hash := sha256.Sum256([]byte(in))
	hashSlice := hash[:6]
	hexString := fmt.Sprintf("%X", hashSlice)
	return hexString
}
