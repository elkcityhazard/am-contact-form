package repository

import (
	"net/http"
	"regexp"
)

// this interface exists so I can pass the router into app config, so I can have access to CtxKey in the handlers

// Route makes up a route for a Router
type Route struct {
	Method  string
	Regex   *regexp.Regexp
	Handler http.HandlerFunc
}

type RouterInterface interface {
	InitRoutes() http.Handler
	NewRoute(method, pattern string, handler http.HandlerFunc) Route
	Serve(w http.ResponseWriter, r *http.Request)
	GetField(r *http.Request, index int) string
}
