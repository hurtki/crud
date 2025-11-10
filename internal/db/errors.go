package db

import (
	"errors"
)

var (
	ErrorNoRowsAffected = errors.New("no rows affected")
)
