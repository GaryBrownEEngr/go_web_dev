package gamestore

import (
	"context"
	"net/http"
	"time"

	"github.com/GaryBrownEEngr/go_web_dev/backend/models"
	"github.com/GaryBrownEEngr/go_web_dev/backend/utils/uerr"
)

type gameStore struct {
	db models.KeyDBStore
}

var _ models.MathGameStore = &gameStore{}

type DbGameData struct {
	Username  string    `dynamodbav:"Name"`
	UpdatedAt time.Time `dynamodbav:"updated_at"`
	JsonData  string    `dynamodbav:"json_data"`
}

func NewUserStore(db models.KeyDBStore) (*gameStore, error) {
	ret := &gameStore{
		db: db,
	}

	return ret, nil
}

func (s *gameStore) Read(username string) (*models.MathGameData, error) {
	dbVal := DbGameData{}
	err := s.db.Get(context.TODO(), username, &dbVal)
	if err != nil {
		return nil, uerr.UErrLogHash("Error reading math game data", http.StatusInternalServerError, err)
	}
	if dbVal.Username == "" || dbVal.Username != username {
		return nil, nil
	}

	ret := models.MathGameData{
		Username:  dbVal.Username,
		UpdatedAt: dbVal.UpdatedAt,
		JsonData:  dbVal.JsonData,
	}

	return &ret, nil
}

func (s *gameStore) Write(in *models.MathGameData) error {
	dbVal := DbGameData{
		Username:  in.Username,
		UpdatedAt: time.Now(),
		JsonData:  in.JsonData,
	}

	err := s.db.Put(context.TODO(), &dbVal)
	if err != nil {
		return uerr.UErrLogHash("Error writing math game data", http.StatusInternalServerError, err)
	}

	return nil
}
