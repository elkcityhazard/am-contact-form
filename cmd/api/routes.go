package main

import (
	"context"
	"net/http"
	"regexp"
	"strings"

	"github.com/elkcityhazard/am-contact-form/internal/handlers"
	"github.com/elkcityhazard/am-contact-form/internal/repository"
)

// Router is a struct that has routes

type Router struct {
	routes []repository.Route
}

// CtxKey handles the request Contexst

type CtxKey struct{}

//	NewRouter returns a new Router with populated routes

func NewRouter() *Router {
	rtr := &Router{}

	rtr.routes = append(rtr.routes, rtr.NewRoute("GET", "/api/v1/healthcheck", handlers.Repo.HandleHealthcheck))
	rtr.routes = append(rtr.routes, rtr.NewRoute("GET", "/api/v1/([^/]+)", RequiresAuth(handlers.Repo.HandleFormSubmission)))

	return rtr
}

//	InitRoutes is an entry point to encapsulate the Router routes with their middleware

func (rtr *Router) InitRoutes() http.Handler {

	return InitMiddleware(http.HandlerFunc(rtr.Serve))
}

// This appends a new route to the slice of routes

func (rtr *Router) NewRoute(method, pattern string, handler http.HandlerFunc) repository.Route {
	return repository.Route{Method: method, Regex: regexp.MustCompile("^" + pattern + "$"), Handler: handler}
}

//	Serve handles the http Request, checking if the route matches, and the method matches

func (rtr *Router) Serve(w http.ResponseWriter, r *http.Request) {

	// populate the allow header if the Method is wrong
	var allow []string

	// 	loop through the routes
	//	check for regex matches
	//	if there are matches, check the method, and append to allow if necessary
	// if the method and route match, create the context key with the matched route param
	//	use the handler for the route

	for _, route := range rtr.routes {

		matches := route.Regex.FindStringSubmatch(r.URL.Path)

		if len(matches) > 0 {
			if r.Method != route.Method {
				allow = append(allow, route.Method)
				continue
			}

			// handle the actual route

			ctx := context.WithValue(r.Context(), CtxKey{}, matches[1:])

			route.Handler(w, r.WithContext(ctx))
			return
		}

	}

	// this means that there was a route match, but the method was wrong

	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.NotFound(w, r)
}

//	GetField takes in the request object, and an index, and returns the CtxKey{} cast to a slice of stinf

func (rtr *Router) GetField(r *http.Request, index int) string {
	fields := r.Context().Value(CtxKey{}).([]string)
	return fields[index]
}
