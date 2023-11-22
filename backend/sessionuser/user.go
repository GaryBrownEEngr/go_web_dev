package sessionuser

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/GaryBrownEEngr/go_web_dev/backend/aws/awsDynamo"
	"github.com/GaryBrownEEngr/go_web_dev/backend/models"
	"github.com/GaryBrownEEngr/go_web_dev/backend/utils"
	"github.com/GaryBrownEEngr/go_web_dev/backend/utils/stacktrs"
	"github.com/GaryBrownEEngr/go_web_dev/backend/utils/uerr"
)

type userStore struct {
	db          models.KeyDBStore
	garbageHash string
}

var _ models.UserStore = &userStore{}

func NewUserStore() (*userStore, error) {
	db, err := awsDynamo.NewDynamoDB("GoWebDev_user", "username")
	if err != nil {
		return nil, stacktrs.Wrap(err)
	}

	garbageHash, err := utils.HashPassword("Garbage Password")
	if err != nil {
		return nil, stacktrs.Wrap(err)
	}

	ret := &userStore{
		db:          db,
		garbageHash: garbageHash,
	}

	return ret, nil
}

type DbUser struct {
	Username       string    `dynamodbav:"username"`
	HashedPassword string    `dynamodbav:"hashed_password"`
	CreatedAt      time.Time `dynamodbav:"created_at"`
}

func (s *userStore) CreateUser(username, password string) (*models.User, error) {
	username = strings.ToLower(username)
	if len(username) < 2 || len(username) > 100 {
		return nil, uerr.UErr("Username is too short", http.StatusBadRequest)
	}
	if len(password) < 6 || len(password) > 100 {
		return nil, uerr.UErr("Password is too short", http.StatusBadRequest)
	}

	// Make it so this function is time constant
	// When the username fails, hash the password anyways.
	// Also alway take at least 1 second.
	// The reason being, if there are enough requests, then hashing could start taking longer than 1 second.
	timer := time.NewTimer(time.Second * 1)
	defer func() {
		<-timer.C
	}()

	hash, err := utils.HashPassword(password)
	if err != nil {
		return nil, uerr.UErrLogHash("Error creating user", http.StatusInternalServerError, err)
	}

	u := DbUser{}
	err = s.db.Get(context.TODO(), username, &u)
	if err != nil {
		return nil, uerr.UErrLogHash("Error creating user", http.StatusInternalServerError, err)
	}
	if u.Username == username {
		return nil, uerr.UErrLog("Username already used: "+username, http.StatusBadRequest, nil)
	}

	u = DbUser{
		Username:       username,
		HashedPassword: hash,
		CreatedAt:      time.Now(),
	}

	err = s.db.Put(context.TODO(), &u)
	if err != nil {
		return nil, uerr.UErrLogHash("Error creating user", http.StatusInternalServerError, err)
	}

	ret := &models.User{
		Username:  username,
		CreatedAt: u.CreatedAt,
	}
	return ret, nil
}

func (s *userStore) VerifyPassword(username, password string) (*models.User, error) {
	username = strings.ToLower(username)
	if len(username) < 2 || len(username) > 100 {
		return nil, uerr.UErr("Username is too short", http.StatusBadRequest)
	}
	if len(password) < 6 || len(password) > 100 {
		return nil, uerr.UErr("Password is too short", http.StatusBadRequest)
	}

	// Make it so this function is time constant
	// When the username fails, hash the password anyways.
	// Also alway take at least 1 second.
	// The reason being, if there are enough requests, then hashing could start taking longer than 1 second.
	timer := time.NewTimer(time.Second * 1)
	hashHasRun := false
	defer func() {
		if !hashHasRun {
			_ = utils.VerifyPassword(s.garbageHash, password)
		}
		<-timer.C
	}()

	u := DbUser{}
	err := s.db.Get(context.TODO(), username, &u)
	if err != nil {
		return nil, uerr.UErrLogHash("Error getting user", http.StatusInternalServerError, err)
	}
	if u.Username == "" {
		return nil, uerr.UErrLogHash("Invalid username or password", http.StatusUnauthorized, fmt.Errorf(username))
	}

	if !utils.VerifyPassword(u.HashedPassword, password) {
		return nil, uerr.UErrLogHash("Invalid username or password", http.StatusUnauthorized, fmt.Errorf(username))
	}

	ret := &models.User{
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
	}

	return ret, nil
}

func (s *userStore) DeleteUser(username string) error {
	err := s.db.Delete(context.TODO(), username)
	if err != nil {
		return uerr.UErrLogHash("Error deleting user", http.StatusInternalServerError, err)
	}

	return nil
}
