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

// TODO: write docs how to write path
func (s *RouteSet) Add(path string, handler http.HandlerFunc) {
	route := NewRoute(path, handler)
	s.routes = append(s.routes, route)
}

func (s *RouteSet) Handler(path string) (http.Handler, any, error) {
	for _, route := range s.routes {

		matches, parameter := route.Match(path)
		if matches {
			return route.handler, parameter, nil
		}
	}

	return nil, nil, ErrNotFound
}
