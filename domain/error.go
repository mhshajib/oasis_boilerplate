package domain

import "errors"

// contains common error for domain
var (
	ErrWriteCache    = errors.New("failed to write cache")
	ErrReadCache     = errors.New("failed to read cache")
	ErrCacheNotFound = errors.New("not found in cache")
	ErrProFeature    = errors.New("limited to pro user")
)
