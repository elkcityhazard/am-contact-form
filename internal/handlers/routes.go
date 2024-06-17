package handlers

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

var activeRouter *Router

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

type Router struct {
	routes []route
}

type CtxKey struct{}

func NewRouter() *Router {
	rtr := &Router{}

	rtr.routes = append(rtr.routes, rtr.NewRoute("GET", "/api/v1/healthcheck", HandlerRepo.HandleHealthcheck))
	rtr.routes = append(rtr.routes, rtr.NewRoute("GET", "/api/v1/([^/]+)", RequiresAuth(HandlerRepo.HandleFormSubmission)))

	activeRouter = rtr

	return rtr
}

func (rtr *Router) InitRoutes() http.Handler {

	return InitMiddleware(http.HandlerFunc(rtr.Serve))
}

func (rtr *Router) NewRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

func (rtr *Router) Serve(w http.ResponseWriter, r *http.Request) {
	var allow []string

	for _, route := range rtr.routes {

		matches := route.regex.FindStringSubmatch(r.URL.Path)

		fmt.Println(matches)

		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}

			// handle the actual route

			ctx := context.WithValue(r.Context(), CtxKey{}, matches[1:])

			route.handler(w, r.WithContext(ctx))
			return
		}

	}

	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.NotFound(w, r)
}

func (rtr *Router) GetField(r *http.Request, index int) string {
	fields := r.Context().Value(CtxKey{}).([]string)
	return fields[index]
}
