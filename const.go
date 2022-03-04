package markov

import (
	"errors"
)

var (
	ErrIdxOutOfRange = errors.New("index out of range")
	ErrInvalidStatus = errors.New("invalid status")
	ErrInvalidRates  = errors.New("invalid rates")
)

const (
	StatusNotAvailable = -1
)
