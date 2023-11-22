package user

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
	"golang.org/x/crypto/bcrypt"
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

	garbageHash, err := hashPassword("Garbage Password")
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
		return nil, utils.NewUserErr("Username is too short", http.StatusBadRequest)
	}
	if len(password) < 6 || len(password) > 100 {
		return nil, utils.NewUserErr("Password is too short", http.StatusBadRequest)
	}

	// Make it so this function is time constant
	// When the username fails, hash the password anyways.
	// Also alway take at least 1 second.
	// The reason being, if there are enough requests, then hashing could start taking longer than 1 second.
	timer := time.NewTimer(time.Second * 1)
	defer func() {
		<-timer.C
	}()

	hash, err := hashPassword(password)
	if err != nil {
		return nil, utils.NewUserErrLogHash("Error creating user", http.StatusInternalServerError, err)
	}

	u := DbUser{}
	err = s.db.Get(context.TODO(), username, &u)
	if err != nil {
		return nil, utils.NewUserErrLogHash("Error creating user", http.StatusInternalServerError, err)
	}
	if u.Username == username {
		return nil, utils.NewUserErrLog("Username already used: "+username, http.StatusBadRequest, nil)
	}

	u = DbUser{
		Username:       username,
		HashedPassword: hash,
		CreatedAt:      time.Now(),
	}

	err = s.db.Put(context.TODO(), &u)
	if err != nil {
		return nil, utils.NewUserErrLogHash("Error creating user", http.StatusInternalServerError, err)
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
		return nil, utils.NewUserErr("Username is too short", http.StatusBadRequest)
	}
	if len(password) < 6 || len(password) > 100 {
		return nil, utils.NewUserErr("Password is too short", http.StatusBadRequest)
	}

	// Make it so this function is time constant
	// When the username fails, hash the password anyways.
	// Also alway take at least 1 second.
	// The reason being, if there are enough requests, then hashing could start taking longer than 1 second.
	timer := time.NewTimer(time.Second * 1)
	hashHasRun := false
	defer func() {
		if !hashHasRun {
			_ = doPasswordsMatch(s.garbageHash, password)
		}
		<-timer.C
	}()

	u := DbUser{}
	err := s.db.Get(context.TODO(), username, &u)
	if err != nil {
		return nil, utils.NewUserErrLogHash("Error getting user", http.StatusInternalServerError, err)
	}
	if u.Username == "" {
		return nil, utils.NewUserErrLogHash("Invalid username or password", http.StatusUnauthorized, fmt.Errorf(username))
	}

	if !doPasswordsMatch(u.HashedPassword, password) {
		return nil, utils.NewUserErrLogHash("Invalid username or password", http.StatusUnauthorized, fmt.Errorf(username))
	}

	ret := &models.User{
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
	}

	return ret, nil
}

func (s *userStore) DeleteUser(in models.User) error {
	return utils.NewUserErrLog("Error deleting user", http.StatusInternalServerError, stacktrs.Errorf("Not implemented"))
}

// Hash password using Bcrypt
func hashPassword(password string) (string, error) {
	// Convert password string to byte slice
	passwordBytes := []byte(password)
	// Hash password with Bcrypt's min cost
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	return string(hashedPasswordBytes), err
}

// Check if two passwords match using Bcrypt's CompareHashAndPassword
func doPasswordsMatch(hashedPassword, currPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(currPassword))
	return err == nil
}
