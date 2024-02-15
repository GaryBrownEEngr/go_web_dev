package models

import (
	"context"
	"errors"
)

var KeyDbNoDocFound = errors.New("Key DB-Document not found")

type KeyDBStore interface {
	Get(ctx context.Context, key string, out interface{}) error
	Put(ctx context.Context, in interface{}) error
	Delete(ctx context.Context, key string) error
}
