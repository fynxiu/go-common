package store

/**
	Deprecated
 */

import "time"

type Store interface {
	Set(key string, value string, expireIn time.Duration) error
	Get(key string) (string, error)
}
