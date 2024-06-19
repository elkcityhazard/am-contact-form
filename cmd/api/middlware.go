package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

func StripTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		switch true {
		case r.URL.Path == "/":
			next.ServeHTTP(w, r.WithContext(r.Context()))
		case r.URL.Path[len(r.URL.Path)-1:] == "/":
			redirectURL := r.URL.Path[0 : len(r.URL.Path)-1]
			http.Redirect(w, r.WithContext(r.Context()), redirectURL, http.StatusPermanentRedirect)
		default:
			next.ServeHTTP(w, r.WithContext(r.Context()))

		}

	})
}

func AddDefaultHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "GET" {
			w.Header().Set("Cache-Control", "max-age=3600")
		}

		w.Header().Set("X-CSRF-TOKEN", nosurf.Token(r))

		next.ServeHTTP(w, r.WithContext(r.Context()))

	})
}

func InitMiddleware(next http.Handler) http.Handler {
	return LoadAndSaveSession(StripTrailingSlash(AddDefaultHeaders(NoSurf(next))))
}

//	RequiresAuth is for specific routes that we use on the handlers themselves.
//	Note: next has been upgraded to HandlerFunc and we now return a HandlerFunc so we can
//	make sure this conforms to the handler prop of the Router type

func RequiresAuth(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		id := app.SessionManager.GetInt(r.Context(), "id")

		if id == 0 {
			http.Redirect(w, r.WithContext(r.Context()), "/login", http.StatusSeeOther)
		}

		next.ServeHTTP(w, r.WithContext(r.Context()))

	})
}

func LoadAndSaveSession(next http.Handler) http.Handler {

	return app.SessionManager.LoadAndSave(next)

}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.ExemptFunc(func(r *http.Request) bool {
		if r.Method == "POST" && r.URL.Path == "/api/v1/hello-world" {
			return true
		}

		return false
	})

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.IsProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}
