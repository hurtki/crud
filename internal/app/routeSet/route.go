package routeSet

import (
	"net/http"
	"strings"
)

type route struct {
	parts   []routePart
	handler http.HandlerFunc
}

func NewRoute(pattern string, handler http.HandlerFunc) route {
	pattern = strings.Trim(pattern, "/")
	patternParts := strings.Split(pattern, "/")

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

	return route{
		parts:   routeParts,
		handler: handler,
	}
}

func (r *route) Match(path string) (bool, any) {
	path = strings.Trim(path, "/")
	pathParts := strings.Split(path, "/")

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
