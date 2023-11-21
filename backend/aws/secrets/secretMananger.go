package secrets

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/GaryBrownEEngr/twertle_api_dev/backend/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/jszwec/csvutil"
)

type cache struct {
	cache      map[string]string
	filled     bool
	filledTime time.Time
	lock       sync.RWMutex
}

var _ models.SecretStore = &cache{}

func NewSecretManager() (*cache, error) {
	ret := &cache{
		cache: make(map[string]string),
	}

	err := ret.fill()
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *cache) fill() error {
	s.lock.Lock()
	defer s.lock.Unlock()

	secretName := "AppRunner/GoWebDev"
	region := "us-west-2"

	// https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/
	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return fmt.Errorf("Error while loading AWS default config: %w", err)
	}

	// Create Secrets Manager client
	svc := secretsmanager.NewFromConfig(config)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("Error while getting secrets: %w", err)
	}

	if result.SecretString == nil {
		return fmt.Errorf("Secret string is missing")
	}

	type csvRow struct {
		Key   string
		Value string
	}
	var rows []csvRow
	err = csvutil.Unmarshal([]byte(*result.SecretString), &rows)
	if err != nil {
		return fmt.Errorf("Error while unmarshaling secrets: %w", err)
	}

	for _, row := range rows {
		s.cache[row.Key] = row.Value
	}
	s.filled = true
	s.filledTime = time.Now()

	return nil
}

func (s *cache) Get(key string) (string, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	ret, ok := s.cache[key]
	if !ok {
		return "", fmt.Errorf("Secet for key '%s' not found", key)
	}

	return ret, nil
}
