package routeSet

import (
	"net/http"
)

// RouteSet is for storing path to handler map
type RouteSet struct {
	routes []route
}

func NewRouteSet() RouteSet {

	return RouteSet{routes: []route{}}
}

// Add(pattern, handler) creates a new route in routeSet with a pattern
// pattern can contain one of parameters: {num}, {string}
// pattern can contain strict parts: /tasks/, /api/users/
// examples: 'api/users/3', 'post/{string}'
func (s *RouteSet) Add(path string, handler http.HandlerFunc) {
	route := NewRoute(path, handler)
	s.routes = append(s.routes, route)
}

// Handler(path string) tris to match given path specified with Add() routes
// returns a handler, url parameter(if there is no then nil), error (usually ErrNotFound)
func (s *RouteSet) Handler(path string) (http.Handler, any, error) {
	for _, route := range s.routes {

		matches, parameter := route.Match(path)
		if matches {
			return route.handler, parameter, nil
		}
	}

	return nil, nil, ErrNotFound
}
