package routeSet

import (
	"errors"
	"fmt"
)

var (
	ErrTwoParameteresInOneRoute     = errors.New("cannot create a route with two parameteres")
	ErrWrongParameterTypeInGivenURL = errors.New("while matching given part and route part, cannot parse given part to right type")
	ErrNotFound                     = errors.New("Path not found")
)

type ErrorBadRoutePart struct {
	comment string
}

func NewErrorBadRoutePart(comment string) ErrorBadRoutePart {
	return ErrorBadRoutePart{comment: comment}
}

func (e ErrorBadRoutePart) Error() string {
	return fmt.Sprintf("error, while creating new route part occured, comment: %s", e.comment)
}
