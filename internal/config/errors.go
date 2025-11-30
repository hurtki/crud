package config

import "errors"

var (
	ErrNoOriginsWithCors          = errors.New("use_cors is true but cors_origins is empty")
	ErrTasksPerPageSmallerThanOne = errors.New("tasks per page count should be bigger than null")
	ErrShutdownTimeSmallerThanNull = errors.New("time_to_shut_down must be > 0")
	ErrWrongPortSpecified = errors.New("wrong port specified")
)
