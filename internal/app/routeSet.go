package app

import "net/http"

type RouteSet struct {
	routes map[string]http.Handler
}

func NewRouteSet() RouteSet {
	routes := make(map[string]http.Handler)

	return RouteSet{routes: routes}
}

func (s *RouteSet) Add(path string, handler http.Handler) {

	if _, ok := s.routes[path]; ok {
		panic("add already exist path")
	}

	s.routes[path] = handler
}

func (s *RouteSet) Handler(path string) (http.Handler, error) {
	handler, ok := s.routes[path]

	if !ok {
		return nil, ErrNotFound
	}

	return handler, nil
}
