package sessionuser

import (
	"net/http"
	"time"

	"github.com/GaryBrownEEngr/go_web_dev/backend/models"
	"github.com/GaryBrownEEngr/go_web_dev/backend/utils"
	"github.com/GaryBrownEEngr/go_web_dev/backend/utils/stacktrs"
)

func CreateSession(tokenMaker models.TokenMaker, user *models.User) (*models.Token, error) {
	if tokenMaker == nil || user == nil {
		return nil, utils.NewUserErrLogHash("Error creating session", http.StatusInternalServerError, stacktrs.Errorf("nil pointer"))
	}

	token, err := tokenMaker.Create(user.Username, time.Minute*15)
	if err != nil {
		return nil, utils.NewUserErrLogHash("Error creating session", http.StatusInternalServerError, stacktrs.Errorf("nil pointer"))
	}

	return token, nil
}
