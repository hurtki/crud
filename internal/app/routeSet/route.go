package routeSet

import (
	"fmt"
	"net/http"
	"strings"
)

// route stores sample of url route
// and matches it with a handler
type route struct {
	parts   []routePart
	handler http.HandlerFunc
}

// NewRoute(pattern, handler) creates a new route with a pattern
// pattern can contain one of parameters: {num}, {string}
// pattern can contain strict parts: /tasks/, /api/users/
// examples: 'api/users/3', 'post/{string}'
func NewRoute(pattern string, handler http.HandlerFunc) route {
	patternParts := strings.FieldsFunc(pattern, func(r rune) bool {
		return r == '/'
	})

	fmt.Println("creating route with len of patternParts:", len(patternParts))
	fmt.Println("patternParts: ", patternParts)

	routeParts := []routePart{}
	wasParameter := false
	for _, patternPart := range patternParts {
		if patternPart == "" {
			continue
		}

		routePart, err := NewRoutePart(patternPart)
		if err != nil {
			panic(err)
		}
		if wasParameter && !routePart.Strict {
			panic(ErrTwoParameteresInOneRoute)
		}
		if !routePart.Strict {
			wasParameter = true
		}
		routeParts = append(routeParts, routePart)
	}
	fmt.Println("route parts: ", routeParts)
	return route{
		parts:   routeParts,
		handler: handler,
	}
}

// Match(path) checks if given path matches pattern of the route
func (r *route) Match(path string) (bool, any) {
	pathParts := strings.FieldsFunc(path, func(r rune) bool {
		return r == '/'
	})

	var resParameter any = nil

	if len(pathParts) != len(r.parts) {
		return false, nil
	}

	for i, routePart := range r.parts {
		pathPart := pathParts[i]
		if pathPart == "" {
			continue
		}

		ok, parameter := routePart.Compare(pathPart)
		if !ok {
			return false, nil
		}

		if parameter != nil {
			resParameter = parameter
		}
	}

	return true, resParameter
}
