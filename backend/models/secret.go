package models

type SecretStore interface {
	Get(key string) (string, error)
}
