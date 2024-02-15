package utils

import (
	"fmt"
	"sync"
	"time"

	"github.com/GaryBrownEEngr/go_web_dev/backend/models"
	"github.com/jszwec/csvutil"
)

type cache struct {
	cache      map[string]string
	filled     bool
	filledTime time.Time
	lock       sync.RWMutex
}

var _ models.SecretStore = &cache{}

func NewSecretManager(csvData string) (*cache, error) {
	ret := &cache{
		cache: make(map[string]string),
	}

	type csvRow struct {
		Key   string
		Value string
	}
	var rows []csvRow
	err := csvutil.Unmarshal([]byte(csvData), &rows)
	if err != nil {
		return nil, fmt.Errorf("Error while un-marshaling secrets: %w", err)
	}

	for _, row := range rows {
		ret.cache[row.Key] = row.Value
	}
	ret.filled = true
	ret.filledTime = time.Now()

	return ret, nil
}

func (s *cache) Get(key string) (string, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	ret, ok := s.cache[key]
	if !ok {
		return "", fmt.Errorf("Secret for key '%s' not found", key)
	}

	return ret, nil
}
