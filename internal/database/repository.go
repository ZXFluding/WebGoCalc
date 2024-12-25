package repository

import "errors"

var (
	NameNotFound = errors.New("name not found")
	NameExists   = errors.New("name exists")
)
