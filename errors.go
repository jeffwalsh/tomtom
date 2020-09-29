package tomtom

import "errors"

var (
	ErrNoAPIKey           = errors.New("No API Key provided")
	ErrCommandDoesntExist = errors.New("command is not supported by api client")
)
