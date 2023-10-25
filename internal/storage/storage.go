package storage

import "errors"

var (
	ErrURLNotFound      = errors.New("url not found")
	ErrURLNotExists     = errors.New("url not exists")
	ErrURLALreadyExists = errors.New("url already exists")
)
