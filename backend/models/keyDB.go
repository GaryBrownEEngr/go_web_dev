package models

import "context"

type KeyDBStore interface {
	Get(ctx context.Context, key string, out interface{}) error
	Put(ctx context.Context, in interface{}) error
}
